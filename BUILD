load("@bazel_gazelle//:def.bzl", "gazelle")
load("@bazel_gazelle//:def.bzl", "DEFAULT_LANGUAGES", "gazelle_binary")

# gazelle:prefix github.com/broadcastertools/cliprate
# gazelle:build_file_name BUILD.bazel
# gazelle:go_naming_convention go_default_library

gazelle_binary(
    name = "gazelle_binary",
    languages = DEFAULT_LANGUAGES,
    visibility = ["//visibility:public"],
)

gazelle(
    name = "gazelle",
    gazelle = "//:gazelle_binary",
)
