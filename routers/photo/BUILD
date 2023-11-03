load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "photo",
    srcs = ["photo.go"],
    importpath = "github.com/nacl2000/photo_album/routers/photo",
    visibility = ["//visibility:public"],
    deps = [
        "//models/image_model",
        "@com_github_gin_gonic_gin//:gin",
    ],
)
