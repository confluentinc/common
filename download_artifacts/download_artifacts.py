import json
import re
import sys

import requests

from global_config import DAConfig
from nexusURL import NexusURL

VERSION_MATCH = r"<version>(.*?)(|-SNAPSHOT)</version>"
NAME_MATCH = r"<name>(.*?)(|-parent)</name>"
SNAPSHOT_LITERAL = "-SNAPSHOT"
CONFLUENT_LITERAL = "confluent-"
POM_PATH = "pom.xml"

class DownloadArtifacts:
    def __init__(self, cg):
        self.cg = cg

    def get_project_name_version(self):
        try:
            with open(POM_PATH) as f:
                for line in f:
                    if self.cg.version == "":
                        version = re.search(VERSION_MATCH, line)
                        if version is not None:
                            self.cg.version = version.groups()[0]
                    if self.cg.project_name == "":
                        project_name = re.search(NAME_MATCH, line)
                        if project_name is not None:
                            self.cg.project_name = project_name.groups()[0]
        except IOError as e:
            print(e)


    def download(self):
        for file_type, download_path in self.cg.download_path.items():
            try:
                r = requests.get(
                        download_path, auth=(self.cg.username, self.cg.password))
            except requests.exceptions.RequestException as e:
                print(e)
                sys.exit(1)
            package_name = CONFLUENT_LITERAL + self.cg.project_name + "-" +\
                self.cg.version + SNAPSHOT_LITERAL + file_type
            try:
                package = open(package_name, 'wb')
                package.write(r.content)
                package.close()
            except IOError as e:
                print(e)


def main():
    cg = DAConfig
    da = DownloadArtifacts(cg)
    da.get_project_name_version()
    nurl = NexusURL(cg)
    nurl.get_m2_settings()
    nurl.get_download_path()
    da.download()


if __name__ == "__main__":
    main()
