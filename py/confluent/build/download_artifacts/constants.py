from enum import Enum

BASE_URL = "https://nexus.confluent.io/service/rest/v1/search/assets?q="
M2 = "~/.m2/settings.xml"
PACKAGE_LITERAL = "-package"
ARTIFACTID = "&maven.artifactId="
USERNAME_REGEX = r"<username>(.*)</username>"
PASSWORD_REGEX = r"<password>(.*)</password>"
FILENAME_REGEX = r"\d{8}.\d{6}-(\d*)-.*(.zip|.tar.gz)$"

VERSION_REGEX = r"<version>(.*?)(|-SNAPSHOT)</version>"
NAME_REGEX = r"<name>(.*?)(|-parent)</name>"
SNAPSHOT_LITERAL = "-SNAPSHOT"
CONFLUENT_LITERAL = "confluent-"
POM = "pom.xml"