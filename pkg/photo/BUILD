load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "photo",
    srcs = ["photo.go"],
    data = ["//config/photo:photo_config"],
    importpath = "github.com/nacl2000/photo_album/pkg/photo",
    visibility = ["//visibility:public"],
    deps = [
        "//pkg/file",
        "//pkg/path",
    ],
)
