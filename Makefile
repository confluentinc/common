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
DEFAULT_APPLY_PATCHES=yes
DEFAULT_DESTDIR=$(CURDIR)/BUILD/
DEFAULT_PREFIX=/usr
DEFAULT_SYSCONFDIR=/etc/$(PACKAGE_TITLE)
DEFAULT_SKIP_TESTS=no


# Whether we should apply patches. This only makes sense for alternate packaging
# systems that know how to apply patches themselves, e.g. Debian.
ifndef APPLY_PATCHES
APPLY_PATCHES=$(DEFAULT_APPLY_PATCHES)
endif

# Whether we should run tests during the build.
ifndef SKIP_TESTS
SKIP_TESTS=$(DEFAULT_SKIP_TESTS)
endif

# Install directories
ifndef DESTDIR
DESTDIR=$(DEFAULT_DESTDIR)
endif
# For platform-specific packaging you'll want to override this to a normal
# PREFIX like /usr or /usr/local. Using the PACKAGE_NAME here makes the default
# zip/tgz files use a format like:
#   kafka-version-scalaversion/
#     bin/
#     etc/
#     share/kafka/
ifndef PREFIX
PREFIX=$(DEFAULT_PREFIX)
endif

ifndef SYSCONFDIR
SYSCONFDIR:=$(DEFAULT_SYSCONFDIR)
endif
SYSCONFDIR:=$(subst PREFIX,$(PREFIX),$(SYSCONFDIR))

export APPLY_PATCHES
export VERSION
export DESTDIR
export PREFIX
export SYSCONFDIR
export SKIP_TESTS

MVN_SETTINGS=-Dskip.maven.resolver.plugin=true -Dcustom.install.phase=none -Dcustom.deploy.phase=none

all: install


archive: install
	rm -f $(CURDIR)/$(PACKAGE_NAME).tar.gz && cd $(DESTDIR) && tar -czf $(CURDIR)/$(PACKAGE_NAME).tar.gz $(PREFIX)
	rm -f $(CURDIR)/$(PACKAGE_NAME).zip && cd $(DESTDIR) && zip -r $(CURDIR)/$(PACKAGE_NAME).zip $(PREFIX)

apply-patches: $(wildcard patches/*)
ifeq ($(APPLY_PATCHES),yes)
	git reset --hard HEAD
	cat patches/series | xargs -iPATCH bash -c 'patch -p1 < patches/PATCH'
endif

build: apply-patches
ifeq ($(SKIP_TESTS),yes)
	mvn -B -DskipTests=true $(MVN_SETTINGS) install
else
	mvn -B $(MVN_SETTINGS) install
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



export RPM_VERSION=$(shell echo $(VERSION) | sed -e 's/-alpha[0-9]*//' -e 's/-beta[0-9]*//' -e 's/-rc[0-9]*//' -e 's/-SNAPSHOT//' -e 's/-cp[0-9]*//' -e 's/-[0-9]*//')
# Get any -alpha, -beta, -rc piece that we need to put into the Release part of
# the version since RPM versions don't support non-numeric
# characters. Ultimately, for something like 0.8.2-beta, we want to end up with
# Version=0.8.2 Release=0.X.beta
# where X is the RPM release # of 0.8.2-beta (the prefix 0. forces this to be
# considered earlier than any 0.8.2 final releases since those will start with
# Version=0.8.2 Release=1)
export RPM_RELEASE_POSTFIX=$(subst -,,$(subst $(RPM_VERSION),,$(VERSION)))
ifneq ($(RPM_RELEASE_POSTFIX),)
	export RPM_RELEASE_POSTFIX_UNDERSCORE=_$(RPM_RELEASE_POSTFIX)
endif

rpm: RPM_BUILDING/SOURCES/$(PACKAGE_TITLE)-$(RPM_VERSION).tar.gz
	echo "Building the rpm"
	rpmbuild --define="_topdir `pwd`/RPM_BUILDING" -tb $<
	find RPM_BUILDING/{,S}RPMS/ -type f | xargs -n1 -iXXX mv XXX .
	echo
	echo "================================================="
	echo "The rpms have been created and can be found here:"
	@ls -laF $(PACKAGE_TITLE)*rpm
	echo "================================================="

# Unfortunately, because of version naming issues and the way rpmbuild expects
# the paths in the tar file to be named, we need to rearchive the package. So
# instead of depending on archive, this target just uses the unarchived,
# installed version to generate a new archive. Note that we always regenerate
# the symlink because the RPM_VERSION doesn't include all the version info -- it
# can leave of things like -beta, -rc1, etc.
RPM_BUILDING/SOURCES/$(PACKAGE_TITLE)-$(RPM_VERSION).tar.gz: rpm-build-area install $(PACKAGE_TITLE).spec.in RELEASE_$(RPM_VERSION)$(RPM_RELEASE_POSTFIX_UNDERSCORE)
	rm -rf $(PACKAGE_TITLE)-$(RPM_VERSION)
	mkdir $(PACKAGE_TITLE)-$(RPM_VERSION)
	cp -R $(DESTDIR)/* $(PACKAGE_TITLE)-$(RPM_VERSION)
	./create_spec.sh $(PACKAGE_TITLE).spec.in $(PACKAGE_TITLE)-$(RPM_VERSION)/$(PACKAGE_TITLE).spec
	rm -f $@ && tar -czf $@ $(PACKAGE_TITLE)-$(RPM_VERSION)
	rm -rf $(PACKAGE_TITLE)-$(RPM_VERSION)

rpm-build-area: RPM_BUILDING/BUILD RPM_BUILDING/RPMS RPM_BUILDING/SOURCES RPM_BUILDING/SPECS RPM_BUILDING/SRPMS

RPM_BUILDING/%:
	mkdir -p $@

$(PACKAGE_TITLE).spec:
	./create_spec.sh $(PACKAGE_TITLE).spec.in $(PACKAGE_TITLE).spec

RELEASE_%:
	echo 0 > $@
