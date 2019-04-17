import json
import re

import requests

from global_config import DownloadArtifactsConfig
from nexusURL import NexusURL
from constants import VERSION_REGEX, NAME_REGEX, SNAPSHOT_LITERAL, CONFLUENT_LITERAL, POM

class DownloadArtifacts:
    def __init__(self):
        cg = DownloadArtifactsConfig
        nurl = NexusURL(cg)
        self.cg = cg
        self.nurl = nurl
        return

    def get_project_name_version(self):
        '''get project name and version from top pom file'''
        try:
            with open(POM) as f:
                for line in f:
                    if self.cg.version == "":
                        version = re.search(VERSION_REGEX, line)
                        if version is not None:
                            self.cg.version = version.groups()[0]
                    if self.cg.project_name == "":
                        project_name = re.search(NAME_REGEX, line)
                        if project_name is not None:
                            self.cg.project_name = project_name.groups()[0]
        except IOError as e:
            print(e)
        return

    def download_artifact(self):
        '''Download artifacts from provided URL'''
        for artifact in self.cg.artifacts.values():
            try:
                r = requests.get(
                        artifact.get('file_url'), auth=(self.cg.username, self.cg.password))
                package_name = CONFLUENT_LITERAL + self.cg.project_name + "-" +\
                    self.cg.version + SNAPSHOT_LITERAL + artifact.get('file_ext')
                package = open(package_name, 'wb')
                package.write(r.content)
                package.close()
            except requests.exceptions.RequestException as e:
                print(e)
            except IOError as e:
                print(e)
        return
    
    def get_maven_settings(self):
        self.nurl.get_maven_settings()
        return
    
    def get_artifact_uri(self):
        self.nurl.get_artifact_uri()
        return

    def download(self):
        self.get_project_name_version()
        self.get_maven_settings()
        self.get_artifact_uri()
        self.download_artifact()
        return

def main():
    da = DownloadArtifacts()
    da.download()

if __name__ == "__main__":
    main()
