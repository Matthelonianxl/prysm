"""
cargo-raze crate build file.

DO NOT EDIT! Replaced on runs of cargo-raze
"""
package(default_visibility = [
  # Public for visibility by "@raze__crate__version//" targets.
  #
  # Prefer access through "//cargo", which limits external
  # visibility to explicit Cargo.toml dependencies.
  "//visibility:public",
])

licenses([
  "restricted", # "Apache-2.0 OR MIT"
])

load(
    "@io_bazel_rules_rust//rust:rust.bzl",
    "rust_library",
    "rust_binary",
    "rust_test",
)


# Unsupported target "benchmarks" with type "bench" omitted

rust_library(
    name = "tinytemplate",
    crate_root = "src/lib.rs",
    crate_type = "lib",
    edition = "2015",
    srcs = glob(["**/*.rs"]),
    deps = [
        "@raze__serde__1_0_91//:serde",
        "@raze__serde_json__1_0_39//:serde_json",
    ],
    rustc_flags = [
        "--cap-lints=allow",
    ],
    version = "1.0.2",
    crate_features = [
    ],
)

