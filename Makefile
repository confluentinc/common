# Dependencies you'll probably need to install to compile this: make, curl, git,
# zip, unzip, patch, java7-jdk | openjdk-7-jdk, maven.

# Release specifics. Note that some of these (VERSION, DESTDIR)
# are required and passed to create_archive.sh as environment variables. That
# script can also pick up some other settings (PREFIX, SYSCONFDIR) to customize
# layout of the installation.
ifndef VERSION
# Note that this is sensitive to this package's version being the first
# <version> tag in the pom.xml
VERSION=$(shell grep \<version\> pom.xml | head -n 1 | awk -F'>|<' '{ print $$3 }')
endif

export PACKAGE_TITLE=confluent-common
export PACKAGE_NAME=$(PACKAGE_TITLE)-$(VERSION)

# Defaults that are likely to vary by platform. These are cleanly separated so
# it should be easy to maintain altered values on platform-specific branches
# when the values aren't overridden by the script invoking the Makefile

# Install directories
DESTDIR=$(CURDIR)/BUILD/

# For platform-specific packaging you'll want to override this to a normal
# PREFIX like /usr or /usr/local. Using the PACKAGE_NAME here makes the default
# zip/tgz files use a format like:
#   kafka-version-scalaversion/
#     bin/
#     etc/
#     share/kafka/
PREFIX=$(PACKAGE_NAME)

# Whether we should run tests during the build.
SKIP_TESTS=no

SYSCONFDIR=PREFIX/etc/$(PACKAGE_TITLE)

# Whether we pull artifacts from nexus/artifactory
PULL_ARTIFACTS=no

SYSCONFDIR:=$(subst PREFIX,$(PREFIX),$(SYSCONFDIR))

export VERSION
export DESTDIR
export PREFIX
export SYSCONFDIR
export SKIP_TESTS
export PULL_ARTIFACTS

all: install

archive: install
ifeq ($(PULL_ARTIFACTS),no)
	rm -f $(CURDIR)/$(PACKAGE_NAME).tar.gz && cd $(DESTDIR) && tar -czf $(CURDIR)/$(PACKAGE_NAME).tar.gz $(PREFIX)
	rm -f $(CURDIR)/$(PACKAGE_NAME).zip && cd $(DESTDIR) && zip -r $(CURDIR)/$(PACKAGE_NAME).zip $(PREFIX)
endif

build:
ifeq ($(PULL_ARTIFACTS),no)
	python3 ./download_artifacts/download_artifacts.py
endif
ifeq ($(SKIP_TESTS),yes)
	mvn -DskipTests=true install
else
	mvn install
endif

install: build
	./create_archive.sh

clean:
	rm -rf $(DESTDIR)
	rm -rf $(CURDIR)/$(PACKAGE_NAME)*
	rm -rf $(PACKAGE_TITLE)-$(RPM_VERSION)*rpm
	rm -rf RPM_BUILDING

distclean: clean

test:

.PHONY: clean install
