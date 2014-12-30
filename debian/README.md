To build packages, in a VM:

    VERSION=0.1
    # Make sure you've updated debian/changelog with dch -i
    git checkout -b debian-$VERSION debian
    git merge $VERSION
    sudo git-buildpackage -us -uc \
        --git-debian-branch=debian-$VERSION \
        --git-upstream-tag=$VERSION \
        --git-verbose
    # Build output will be in parent directory, e.g ../confluent-common_0.1-1_all.deb
    git checkout master && git branch -D debian-$VERSION

Note that you may need to install some build dependencies (e.g. make, Maven,
git-buildpackage, debuilder, etc.)

If you are testing before a tagged release is made you can use a branch instead:

    git checkout -b debian-$VERSION debian
    git merge master
    sudo git-buildpackage -us -uc \
        --git-debian-branch=debian-$VERSION \
        --git-upstream-tree=branch --git-upstream-branch=master \
        --git-verbose
    git checkout master && git branch -D debian-$VERSION

If you're working on updating the build script, you may also find the option
`--git-ignore-new` handy, but it shouldn't be used for final builds -- make sure
you start from a clean tree in that case.
