load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "login",
    srcs = ["login.go"],
    importpath = "github.com/nacl2000/photo_album/routers/login",
    visibility = ["//visibility:public"],
    deps = [
        "//models/user_model",
        "@com_github_gin_gonic_contrib//sessions",
        "@com_github_gin_gonic_gin//:gin",
    ],
)
