import json
import re
import time

import requests

from constants import DownloadArtifactsConstants as DAC
from global_config import DownloadArtifactsConfig
from nexusURL import NexusURL


class DownloadArtifacts:
    ''' download the newest artifact from Nexus '''
    def __init__(self):
        config = DownloadArtifactsConfig
        nexus_url = NexusURL(config)
        self.config = config
        self.nexus_url = nexus_url
        return

    def get_project_name_version(self):
        ''' get project name and version from top pom file '''
        try:
            with open(DAC.POM.value) as f:
                for line in f:
                    if self.config.version == "":
                        version = re.search(DAC.VERSION_REGEX.value, line)
                        if version is not None:
                            self.config.version = version.groups()[0]
                    if self.config.project_name == "":
                        project_name = re.search(DAC.NAME_REGEX.value, line)
                        if project_name is not None:
                            self.config.project_name = project_name.groups()[0]
        except IOError as e:
            print(e)
            print("Top level pom.xml is not set right")
        return

    def download_artifact(self):
        ''' Download artifacts from provided URL '''
        for artifact in self.config.artifacts.values():
            try:
                session = requests.Session()
                session.trust_env = False
                r = session.get(
                        artifact.get('file_url'), auth=(
                            self.config.username, self.config.password))
                package_name = DAC.CONFLUENT_LITERAL.value + self.config.project_name + "-" +\
                    self.config.version + DAC.SNAPSHOT_LITERAL.value + artifact.get('file_ext')
                package = open(package_name, 'wb')
                package.write(r.content)
                package.close()
            except requests.exceptions.RequestException as e:
                print(e)
                print("error happened during downloading artifacts")
            except IOError as e:
                print(e)
                print("cannot write package to disk")
        return

    def get_maven_settings(self):
        self.nexus_url.get_maven_settings()
        return

    def get_artifact_uri(self):
        self.nexus_url.get_artifact_uri()
        return

    def download(self):
        self.get_project_name_version()
        self.get_maven_settings()
        self.get_artifact_uri()
        self.download_artifact()
        return


def main():
    start_time = time.time()
    dawnload_artifacts = DownloadArtifacts()
    dawnload_artifacts.download()
    print("--- used %s seconds to download artifacts! ---" % (time.time() - start_time))


if __name__ == "__main__":
    main()
