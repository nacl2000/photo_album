load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "routers",
    srcs = [
        "homepage.go",
        "login.go",
        "photo.go",
        "upload.go",
    ],
    importpath = "github.com/nacl2000/photo_album/routers",
    visibility = ["//visibility:public"],
    deps = [
        "//models/user_model",
        "//pkg/auth",
        "//pkg/network",
        "//pkg/photo",
        "@com_github_gin_gonic_contrib//sessions",
        "@com_github_gin_gonic_gin//:gin",
    ],
)
