import os
import re

import requests

from constants import BASE_URL, M2, PACKAGE_LITERAL, ARTIFACTID, USERNAME_REGEX, PASSWORD_REGEX, FILENAME_REGEX


class NexusURL:
    def __init__(self, cg):
        self.cg = cg

    def get_maven_settings(self):
        '''Get maven password and username from local m2.settings'''
        m2_path = os.path.expanduser(M2)
        try:
            with open(m2_path) as f:
                for line in f:
                    if re.search(USERNAME_REGEX, line) is not None and self.cg.username == "":
                        self.cg.username = re.search(USERNAME_REGEX, line).groups()[0]
                    if re.search(PASSWORD_REGEX, line) is not None and self.cg.password == "":
                        self.cg.password = re.search(PASSWORD_REGEX, line).groups()[0]
        except IOError as e:
            print(e)
        return

    def get_artifact_uri(self):
        '''get latest artifact url'''
        latest = 0
        url = BASE_URL + self.cg.project_name + PACKAGE_LITERAL + "-" + self.cg.version + ARTIFACTID + self.cg.project_name + PACKAGE_LITERAL
        try:
            session = requests.Session()
            session.trust_env = False
            r = session.get(url, auth=(self.cg.username, self.cg.password))
        except requests.exceptions.RequestException as e:
            print(e)
        artifacts = r.json().get('items')
        for artifact in artifacts:
            path = artifact.get('path')
            path_split = path.split("-"+self.cg.version+"-", 1)
            if len(path_split) > 1:
                match = re.match(FILENAME_REGEX, path_split[1])
                if match is not None:
                    if int(match.groups()[0]) >= latest:
                        latest = int(match.groups()[0])
                        file_type = (match.groups()[1])[1:]
                        self.cg.artifacts.get(file_type)['file_url'] = artifact.get('downloadUrl')
        return
