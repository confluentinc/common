import os
import re
import sys

import requests

BASE_URL = "https://nexus.confluent.io/service/rest/v1/search/assets?q="
M2_PATH = "~/.m2/settings.xml"
PACKAGE_LITERAL = "-package"
ARTIFACTID = "&maven.artifactId="
USERNAME_MATCH = r"<username>(.*)</username>"
PASSWORD_MATCH = r"<password>(.*)</password>"
FILENAME_MATCH = r"\d{8}.\d{6}-(\d*)-.*(.zip|.tar.gz)$"


class NexusURL:
    def __init__(self, cg):
        self.cg = cg

    def get_m2_settings(self):
        m2_path = os.path.expanduser(M2_PATH)
        with open(m2_path) as f:
            for line in f:
                if re.search(USERNAME_MATCH, line) is not None and self.cg.username == "":
                    self.cg.username = re.search(USERNAME_MATCH, line).groups()[0]
                if re.search(PASSWORD_MATCH, line) is not None and self.cg.password == "":
                    self.cg.password = re.search(PASSWORD_MATCH, line).groups()[0]

    def get_download_path(self):
        latest = 0
        url = BASE_URL + self.cg.project_name + PACKAGE_LITERAL + "-" + self.cg.version + ARTIFACTID + self.cg.project_name + PACKAGE_LITERAL
        try:
            r = requests.get(url, auth=(self.cg.username, self.cg.password))
        except requests.exceptions.RequestException as e:
            print(e)
            sys.exit(1)
        artifacts = r.json().get('items')
        for artifact in artifacts:
            path = artifact.get('path')
            path_split = path.split("-"+self.cg.version+"-", 1)
            if len(path_split) > 1:
                match = re.match(FILENAME_MATCH, path_split[1])
                if match is not None:
                    if int(match.groups()[0]) >= latest:
                        latest = int(match.groups()[0])
                        self.cg.download_path[match.groups()[1]] = artifact.get('downloadUrl')
        print(self.cg.download_path)
