import os
import re

import requests

from constants import NexusConstants as NC


class NexusURL:
    ''' find the right artifacts to download from nexus '''
    def __init__(self, nexus_config):
        self.nexus_config = nexus_config

    def get_maven_settings(self):
        '''Get maven password and username from local m2.settings'''
        m2_path = os.path.expanduser(NC.MVN_SETTINGS.value)
        try:
            with open(m2_path) as f:
                for line in f:
                    if re.search(NC.USERNAME_REGEX.value, line) is not None\
                            and self.nexus_config.username == "":
                        self.nexus_config.username = re.search(
                            NC.USERNAME_REGEX.value, line).groups()[0]
                    if re.search(NC.PASSWORD_REGEX.value, line) is not None\
                            and self.nexus_config.password == "":
                        self.nexus_config.password = re.search(
                            NC.PASSWORD_REGEX.value, line).groups()[0]
        except IOError as e:
            print(e)
            print("MVN settings is not set right")
        return

    def get_artifact_uri(self):
        '''get latest artifact url'''
        latest = 0
        url = NC.BASE_URL.value + self.nexus_config.project_name\
            + NC.PACKAGE_LITERAL.value + "-" + self.nexus_config.version\
            + NC.ARTIFACTID.value + self.nexus_config.project_name\
            + NC.PACKAGE_LITERAL.value
        try:
            session = requests.Session()
            session.trust_env = False
            r = session.get(url, auth=(
                self.nexus_config.username, self.nexus_config.password))
        except requests.exceptions.RequestException as e:
            print(e)
            print("cannot access to NEXUS, check your MVN setting")
        artifacts = r.json().get('items')
        for artifact in artifacts:
            path = artifact.get('path')
            path_split = path.split("-"+self.nexus_config.version+"-", 1)
            if len(path_split) > 1:
                match = re.match(NC.FILENAME_REGEX.value, path_split[1])
                if match is not None:
                    if int(match.groups()[0]) >= latest:
                        latest = int(match.groups()[0])
                        file_type = (match.groups()[1])[1:]
                        self.nexus_config.artifacts.get(file_type)['file_url']\
                            = artifact.get('downloadUrl')
        return
