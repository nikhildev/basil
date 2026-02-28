# Setup instuctions

This repo needs the following to be installed

- Go 1.26.0 `brew install go`
- Bazelisk 7.3.1 `brew install bazelisk`
  - Note that Bazelisk has Bazel included, so you don't have to install Bazel separately
- Go tools `go get golang.org/x/tools`
  - Make sure to run this after you've installed Go
- [Visual Studio Code with the Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go)
- [Bazel extension (recommended)](https://marketplace.visualstudio.com/items?itemName=BazelBuild.vscode-bazel)
- [Prettier code formatter (recommended)](https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode)

At some point I will probably create a devcontainer with all these installed. But until then we need to install the above manually.