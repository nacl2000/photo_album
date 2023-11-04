load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "homepage",
    srcs = ["homepage.go"],
    importpath = "github.com/nacl2000/photo_album/routers/homepage",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/network",
        "@com_github_gin_gonic_contrib//sessions",
        "@com_github_gin_gonic_gin//:gin",
    ],
)
