# Local Java build wrappers for a Confluent Bazel repo.
#
# Copy this to `tools-bazel/library_wrappers.bzl` in the target repo (drop the
# `.tmpl` suffix) and load its symbols from your BUILD files instead of the raw
# rules. Distilled from ce-kafka's tools-bazel/library_wrappers.bzl — Java only,
# trimmed to the macros common actually uses (java_export + the JUnit5 suite).
#
# Why wrappers? They give every module the same org defaults in one place:
#   - a single POM template for published artifacts,
#   - the `excluded_workspaces` set so tool jars never leak into published jars,
#   - the confluent_rules JUnit5 test-suite macro under a stable local name.
#
# Loads assume: bazel_dep on contrib_rules_jvm and the confluent_rules
# archive_override (see MODULE.bazel from Phase 1).

load("@confluent_rules//java:java_junit_test_suite.bzl", "java_junit_test_suite")
load("@contrib_rules_jvm//java:defs.bzl", oss_java_export = "java_export")

# Tool/proto/grpc workspaces whose jars must never end up inside a published
# artifact. Extend per repo as strict-deps/packaging surfaces more.
_JAR_PKG_EXCLUDED_WORKSPACES = {
    "protobuf": None,
    "com_google_protobuf": None,
    "contrib_rules_jvm_deps": None,
}

# Default POM template used for every published artifact. Copy assets/pom.tpl to
# //tools-bazel:pom.tpl. Preserves io.confluent coordinates at $(maven_version).
_POM_TEMPLATE = "//tools-bazel:pom.tpl"

def java_export(name, pom_template = _POM_TEMPLATE, exclusions = {}, **attrs):
    """Publishable library. Wraps contrib_rules_jvm java_export.

    Generates `<name>.publish` (a maven_publish target) automatically. Pass
    `maven_coordinates = "io.confluent:<artifact>:$(maven_version)"` so the CP
    nano-version flows into the coordinates (see Phase 5). Prefer this for any
    module that other repos consume from CodeArtifact.
    """
    oss_java_export(
        name = name,
        pom_template = pom_template,
        exclusions = exclusions,
        excluded_workspaces = _JAR_PKG_EXCLUDED_WORKSPACES,
        **attrs
    )

def jvm_junit_test_suite(**attrs):
    """JUnit5 test suite: one java_junit5_test per `*Test.java`, plus a root
    test_suite. Thin alias over confluent_rules so BUILD files use a stable
    local name. See Phase 4 for required runtime_deps.
    """
    java_junit_test_suite(**attrs)
