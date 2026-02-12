#!/usr/bin/python

import argparse
import os
import logging
import re
import shlex
import subprocess
import sys

from xml.etree import ElementTree as ET

logging.basicConfig(level=logging.INFO, format='%(message)s')
log = logging.getLogger(__name__)

NAME_SPACE = "{http://maven.apache.org/POM/4.0.0}"
RESOLVER_PLUGIN_VERSION = "0.6.0"


class CI:

    def __init__(self, new_version, repo_path, direct_pom_edit=False):
        """Initialize class variables"""
        # List of all the files that were modified by this script so the parent
        # script that runs this update can commit them.
        self.updated_files = []
        # The new version
        self.new_version = new_version
        # The path the root of the repo so we can use full absolute paths
        self.repo_path = repo_path
        # Use direct pom.xml editing instead of mvn versions:set-property.
        self.direct_pom_edit = direct_pom_edit

    def run_update(self):
        """Update all the files with the new version"""
        log.info("Running additional version updates for common")
        self.resolve_ccs_kafka_version()
        self.resolve_ce_kafka_version()
        self.updated_files.append("pom.xml")
        log.info("Finished all common additional version updates.")

    def resolve_ccs_kafka_version(self):
        """Resolve the version range property for ccs kafka."""
        log.info("Resolving the version range for ccs kafka.")
        property_name="kafka.version"
        version_range = self.get_version_range(property_name)
        version = self.resolve_version_range(version_range, "CCS")
        self.set_property(property_name=property_name, property_value=version)
        log.info("Finished resolving the version range for ccs kafka.")

    def resolve_ce_kafka_version(self):
        """Resolve the version range property for ce kafka."""
        log.info("Resolving the version range for ce kafka.")
        property_name="ce.kafka.version"
        version_range = self.get_version_range(property_name)
        version = self.resolve_version_range(version_range, "CE")
        self.set_property(property_name=property_name, property_value=version)
        log.info("Finished resolving the version range for ce kafka.")

    def get_version_range(self, property_name):
        """Parse pom file and extract property value."""
        log.info("Getting version range for: {}.".format(property_name))
        pom = ET.ElementTree()
        pom.parse(os.path.join(self.repo_path, "pom.xml"))
        properties = pom.getroot().find("{}properties".format(NAME_SPACE))
        version_range = properties.find("{}{}".format(NAME_SPACE, property_name)).text

        if version_range is not None:
            version_range = version_range.strip()
            log.info("Version range for {} is: {}".format(property_name, version_range))
            return version_range
        else:
            log.error("Failed to get value for property: {}".format(property_name))
            sys.exit(1)

    def resolve_version_range(self, version_range, print_method):
        """Run the custom maven resolver plugin to find the latest artifact version in the range."""
        if not self.is_version_range(version_range):
            log.info("Specified version is not a range: {}".format(version_range))
            return version_range

        # We just use one of the kafka artifacts to resolve the range. No particular reason for using this artifact.
        group_id = "org.apache.kafka"
        artifact_id = "kafka-clients"
        cmd = "mvn --batch-mode -Pjenkins "
        # When running this from packaging we need to provide the maven command additional args,
        # such as point at another repo. In packaging if we don't override the repo then it
        # doesn't find the resolver plugin.
        mvn_args = os.getenv("MAVEN_ARGS")

        if mvn_args:
            cmd += "{} ".format(mvn_args)

        cmd += "io.confluent:resolver-maven-plugin:{}:resolve-kafka-range ".format(RESOLVER_PLUGIN_VERSION)
        cmd += "-DgroupId={} ".format(group_id)
        cmd += "-DartifactId={} ".format(artifact_id)
        cmd += "-DversionRange=\"{}\" ".format(version_range)
        # If we don't set this property to false it will skip running this when
        # run from Jenkins because the profile sets the skip value to true.
        cmd += "-Dskip.maven.resolver.plugin=false "
        cmd += "-Dprint{} -q".format(print_method)
        log.info("Resolving the version range for: {}".format(version_range))
        result, stdout = self.run_cmd(cmd, return_stdout=True)

        if result:
            # When run from Jenkins there will be additional output included so we just get the last line of output.
            version = stdout.strip().splitlines()[-1]
            # Make sure we got a valid version back and the script didn't silently fail.
            match_result = re.match("[0-9]*\\.[0-9]*\\.[0-9]*-[0-9]*-{}".format(print_method.lower()), version)

            if match_result:
                log.info("Resolved the version range to version: {}".format(version))
                return version

        log.error("Failed to resolve the version range.")

    def is_version_range(self, version):
        """Checks if the specified Maven version is a range"""
        return (version.startswith('[') or version.startswith('(')) and \
            (version.endswith(']') or version.endswith(')'))

    def run_cmd(self, cmd, return_stdout=False):
        """Execute a shell command. Return true if successful, false otherwise."""
        log.info(cmd)
        cmd = shlex.split(cmd)
        proc = subprocess.Popen(cmd,
                                cwd=self.repo_path,
                                stdout=subprocess.PIPE,
                                stderr=subprocess.STDOUT,
                                universal_newlines=True)
        stdout, _ = proc.communicate()

        if stdout:
            log.info(stdout)

        if proc.returncode != 0:
            if return_stdout:
                return False, stdout
            else:
                return False
        elif return_stdout:
            return True, stdout
        else:
            return True

    def set_property(self, property_name, property_value):
        """Update the project version property in the pom file.

        When --direct-pom-edit is passed to ci-tools ci-update-version, this
        method directly edits the XML instead of invoking mvn versions:set-property
        (which is extremely slow on multi-module projects due to dependency resolution).
        """
        log.info("Setting the property {} to {}".format(property_name, property_value))

        if self.direct_pom_edit:
            self._set_property_direct(property_name, property_value)
        else:
            self._set_property_mvn(property_name, property_value)

        log.info("Finished setting the property.")

    def _set_property_mvn(self, property_name, property_value):
        """Update a pom property using mvn versions:set-property."""
        cmd = "mvn --batch-mode versions:set-property "
        cmd += "-DgenerateBackupPoms=false "
        cmd += "-Dproperty={} ".format(property_name)
        cmd += "-DnewVersion={}".format(property_value)

        if not self.run_cmd(cmd):
            log.error("Failed to set the property.")
            sys.exit(1)

    def _set_property_direct(self, property_name, property_value):
        """Update a pom property by directly editing the XML file."""
        pom_path = os.path.join(self.repo_path, "pom.xml")

        with open(pom_path, 'r') as f:
            content = f.read()

        pattern = r'(<{0}>)[^<]*(</{0}>)'.format(re.escape(property_name))
        new_content, count = re.subn(
            pattern,
            lambda m: m.group(1) + property_value + m.group(2),
            content
        )

        if count == 0:
            log.error("Failed to find property {} in pom.xml".format(property_name))
            sys.exit(1)

        with open(pom_path, 'w') as f:
            f.write(new_content)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description="Run common ci update for nano versioning.")
    parser.add_argument("version", help="The new version to use.")
    parser.add_argument("path", help="The path to the repo.")
    args = parser.parse_args()
    updater = CI(args.version, args.path)
    updater.run_update()
