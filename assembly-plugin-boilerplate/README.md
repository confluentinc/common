# Assembly Plugin Docker Boilerplate

This project uses the assembly plugin rules from resource.xml to package
an artifact containing the raw contents of package.xml. See src/assembly,
at the top level of this repo, for these files.

common includes this subproject as a dependency. Child projects may
simply depend on common-docker, instead of recreating the same boilerplate
XML file all over the place. If common is overkill, projects may depend
on assembly-plugin-boilerplate directly.
