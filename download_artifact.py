import json
import re
import os.path
from os.path import expanduser

import requests

import config
def get_project_name_version(Config):
    with open("pom.xml") as f:
        for line in f:
            if Config.version=="":
                version = re.search(r"<version>(.*?)(|-SNAPSHOT)</version>", line)
                if version != None:
                    Config.version = version.groups()[0]
                    print(Config.version)
            if Config.project_name=="":
                project_name = re.search(r"<name>(.*?)(|-parent)</name>", line)
                if project_name != None:
                    Config.project_name = project_name.groups()[0]
                    print(Config.project_name)

def get_m2_settings(Config):
    m2_path = os.path.expanduser("~/.m2/settings.xml")
    with open(m2_path) as f:
        for line in f:
            if re.search(r"<username>(.*)</username>", line) != None and Config.username=="":
                Config.username = re.search(r"<username>(.*)</username>", line).groups()[0]
            if re.search(r"<password>(.*)</password>", line) != None and Config.password=="":
                Config.password = re.search(r"<password>(.*)</password>", line).groups()[0]

def get_download_path(Config):
    latest = 0
    url = "https://nexus.confluent.io/service/rest/v1/search/assets?q="+Config.project_name+"-package-"+Config.version+"&maven.artifactId="+Config.project_name+"-package"
    r = requests.get(url,auth=(Config.username,Config.password))
    artifacts = r.json().get('items')
    for artifact in artifacts:
        path = artifact.get('path')
        path_split = path.split("-"+Config.version+"-",1)
        if len(path_split) > 1:
            match = re.match(r"\d{8}.\d{6}-(\d*)-.*(.zip|.tar.gz)$",path_split[1])
            if match is not None:
                if int(match.groups()[0]) >= latest:
                    latest = int(match.groups()[0])
                    Config.download_path[match.groups()[1]] = artifact.get('downloadUrl')
    print(Config.download_path)

def download(Config):
    for formate,download_path in Config.download_path.items():
        r = requests.get(download_path,auth=(Config.username,Config.password))
        package_name = "confluent-"+Config.project_name+"-"+Config.version+"-SNAPSHOT"+formate
        package = open(package_name,'wb')
        package.write(r.content)
        package.close()

def main():
    cg = config.TestConfig
    get_project_name_version(cg)
    get_m2_settings(cg)
    get_download_path(cg)
    download(cg)

if __name__ == "__main__":
    main()
