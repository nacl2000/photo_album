load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "auth",
    srcs = ["auth.go"],
    importpath = "github.com/nacl2000/photo_album/pkg/auth",
    visibility = ["//visibility:public"],
    deps = [
        "@com_github_gin_gonic_contrib//sessions",
        "@com_github_gin_gonic_gin//:gin",
    ],
)
