![Bazel logo](https://github.com/user-attachments/assets/13cd50ca-b95e-4fde-b1b5-f94a22618488)

# Basil

A reference implementation for building Go applications in a monorepo using **Bazel**, **Gazelle**, and **bzlmod**.

You can read the full intro on my blog: [Building with Go in a Monorepo using Bazel, Gazelle, and bzlmod](https://nixclix.com/building-with-go-in-a-monorepo-using-bazel-gazelle-and-bzlmod/)

## Project Structure

```
basil/
├── libraries/          # Shared, reusable library packages
│   └── humanize_filesize/
├── packages/           # Standalone application packages
│   └── helloworld/
├── services/           # Microservices that depend on libraries
│   └── service1/
├── MODULE.bazel        # Bazel module definition (bzlmod)
├── BUILD.bazel         # Root build config with Gazelle rules
├── go.mod              # Go module definition
└── MAKEFILE            # Convenience build commands
```

- **libraries/** — Shared code consumed by services and packages. Published with `//visibility:public`.
- **packages/** — Self-contained binaries (e.g., CLI tools, demos).
- **services/** — Applications that compose libraries into runnable services.

## Quick Start

See [SETUP.md](SETUP.md) for installation prerequisites.

```bash
# Generate/update BUILD files from Go source
make gazelle

# Build all targets
make build

# Run tests
make test

# Run individual targets
make run-helloworld
make run-service1
```

## Adding a New Package

1. Create a directory under `libraries/`, `packages/`, or `services/`.
2. Write your Go source files.
3. Run `make gazelle` to auto-generate the `BUILD.bazel` file.
4. Run `make build` to verify.

## Tech Stack

- **Go 1.26.0**
- **Bazel 8.1.0** (via Bazelisk)
- **Gazelle 0.47.0** — automatic BUILD file generation
- **rules_go 0.60.0** — Go support for Bazel
- **bzlmod** — Bazel's module system for dependency management
