<?xml version="1.0" encoding="UTF-8"?>
<!-- POM template for published Confluent Java artifacts (rules_jvm_external
     java_export/maven_export pom_template). Copy to //tools-bazel:pom.tpl.
     {groupId}/{artifactId}/{version}/{dependencies} are substituted by the
     build; {version} resolves to $(maven_version) (the CP nano-version). -->
<project xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd"
         xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <modelVersion>4.0.0</modelVersion>

    <groupId>{groupId}</groupId>
    <artifactId>{artifactId}</artifactId>
    <version>{version}</version>
    <packaging>jar</packaging>
    <url>https://www.confluent.io/</url>
    <licenses>
        <license>
            <name>Apache License 2.0</name>
            <url>http://www.apache.org/licenses/LICENSE-2.0.html</url>
            <distribution>repo</distribution>
        </license>
    </licenses>

    <dependencies>
{dependencies}
    </dependencies>
</project>
