# Local Java build wrappers for a Confluent Bazel repo.
#
# Copy this to `tools-bazel/library_wrappers.bzl` in the target repo (drop the
# `.tmpl` suffix) and load its symbols from your BUILD files instead of the raw
# rules. Distilled from ce-kafka's tools-bazel/library_wrappers.bzl — Java only.
#
# Why wrappers? They give every module the same org defaults in one place:
#   - a single POM template for published artifacts,
#   - a `-neverlink` twin for compile-only ("provided"/Gradle compileOnly) deps,
#   - the `excluded_workspaces` set so tool jars never leak into published jars,
#   - the confluent_rules JUnit5 test-suite macro under a stable local name.
#
# Loads assume: bazel_dep on rules_java, rules_jvm_external, contrib_rules_jvm,
# and the confluent_rules archive_override (see MODULE.bazel from Phase 1).

load("@confluent_rules//java:java_junit_test_suite.bzl", "java_junit_test_suite")
load("@contrib_rules_jvm//java:defs.bzl", oss_java_export = "java_export")
load("@rules_java//java:defs.bzl", oss_java_library = "java_library", oss_java_test = "java_test")
load("@rules_jvm_external//:defs.bzl", oss_maven_export = "maven_export")

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

def java_library(name, neverlink_option = False, **attrs):
    """java_library with an optional compile-only (`-neverlink`) twin.

    Set neverlink_option = True to also emit `<name>-neverlink` (Gradle
    `compileOnly` / Maven `provided` equivalent): compiled against, not shipped.
    """
    oss_java_library(name = name, **attrs)

    if neverlink_option:
        oss_java_library(
            name = name + "-neverlink",
            neverlink = True,
            tags = ["maven:compile-only"],
            exports = [name],
            visibility = attrs.get("visibility", ["//visibility:private"]),
        )

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

def maven_export(name, lib_name, pom_template = _POM_TEMPLATE, exclusions = {}, **attrs):
    """Attach publishing to an existing java_library (the cc-unified-storage
    pattern). Use when a module must stay a plain java_library that other
    in-repo targets depend on, and you add publishing separately.

    The java_library carries `tags = ["maven_coordinates=io.confluent:<art>:$(maven_version)"]`;
    this target references it via `lib_name`. Other maven_exports that depend on
    this artifact must depend on THIS target (not the java_library) to get a
    correctly structured POM.
    """
    oss_maven_export(
        name = name,
        lib_name = lib_name,
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

def _fqcn(src, package_root):
    """"src/test/java/io/confluent/foo/BarTest.java" + root "src/test/java"
    -> "io.confluent.foo.BarTest"."""
    marker = package_root.rstrip("/") + "/"
    idx = src.find(marker)
    rel = src[idx + len(marker):] if idx >= 0 else src
    return rel[:-len(".java")].replace("/", ".")

def jvm_testng_test_suite(
        name,
        srcs,
        deps,
        runtime_deps = [],
        package_root = "src/test/java",
        test_suffixes = ["Test.java", "Tests.java", "IT.java"],
        resources = [],
        jvm_flags = [],
        data = [],
        size = "medium",
        tags = [],
        visibility = None,
        **kwargs):
    """TestNG test suite. contrib_rules_jvm has NO TestNG runner, so this drives
    TestNG through a thin launcher (//tools-bazel/testng:bazel_testng_runner — copy
    assets/BazelTestNgRunner.java + its BUILD alongside this file):

      java_test(use_testrunner = False,
                main_class = "tools.bazel.testng.BazelTestNgRunner",
                args = [<FQCN>])

    Why the launcher, not main_class="org.testng.TestNG": TestNG's own main exit code
    is a status BITMASK (failure|skip|fsp) stricter than maven-surefire — a
    RetryAnalyzer class (fail-then-pass) leaves skip/retry bits set and exits non-zero
    even though the module is green under Maven. The launcher exits non-zero only on a
    genuine hasFailure(), matching surefire.

    One java_test per TestNG `*Test.java` (mirrors jvm_junit_test_suite granularity);
    the module's test sources compile ONCE into `<name>_lib`. For a module that MIXES
    TestNG with JUnit5, pass only the TestNG `*Test.java` here and the JUnit5 ones to
    jvm_junit_test_suite; both can share the compiled test lib.

    Tests that bind a fixed port / machine-global resource MUST pass
    tags = ["exclusive"]: Bazel runs each java_test as its OWN parallel process, so
    Maven's single-JVM in-process coordination (PortMutex etc.) does not hold and
    fixed ports collide; `exclusive` serializes them. See Phase 4.
    """
    lib = name + "_lib"
    oss_java_library(
        name = lib,
        srcs = srcs,
        deps = deps,
        resources = resources,
        testonly = True,
        visibility = ["//visibility:private"],
    )
    rt = [
        ":" + lib,
        "//tools-bazel/testng:bazel_testng_runner",
        "@maven//:org_testng_testng",
    ] + runtime_deps
    test_fqcns = [
        _fqcn(src, package_root)
        for src in srcs
        if [s for s in test_suffixes if src.endswith(s)]
    ]

    # Target name = simple class name when unique; fall back to the fully-qualified
    # name (dots -> underscores) only for collisions, so two classes with the same
    # simple name in different packages (com.a.FooTest / com.b.FooTest) don't emit
    # duplicate `:FooTest` targets (a load-time error).
    simple_counts = {}
    for fqcn in test_fqcns:
        simple = fqcn.split(".")[-1]
        simple_counts[simple] = simple_counts.get(simple, 0) + 1

    tests = []
    for fqcn in test_fqcns:
        simple = fqcn.split(".")[-1]
        tname = simple if simple_counts[simple] == 1 else fqcn.replace(".", "_")
        oss_java_test(
            name = tname,
            use_testrunner = False,
            main_class = "tools.bazel.testng.BazelTestNgRunner",
            args = [fqcn],
            runtime_deps = rt,
            jvm_flags = jvm_flags,
            data = data,
            size = size,
            tags = tags,
            visibility = visibility,
            **kwargs
        )
        tests.append(":" + tname)
    native.test_suite(name = name, tests = tests, visibility = visibility)
