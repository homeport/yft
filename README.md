# ytf [![License](https://img.shields.io/github/license/homeport/yft.svg)](https://github.com/homeport/yft/blob/main/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/homeport/yft)](https://goreportcard.com/report/github.com/homeport/yft) [![Build and Tests](https://github.com/homeport/yft/workflows/Build%20and%20Tests/badge.svg)](https://github.com/homeport/yft/actions?query=workflow%3A%22Build+and+Tests%22) [![Codecov](https://img.shields.io/codecov/c/github/homeport/yft/main.svg)](https://codecov.io/gh/homeport/yft) [![Go Reference](https://pkg.go.dev/badge/github.com/homeport/yft.svg)](https://pkg.go.dev/github.com/homeport/yft) [![Release](https://img.shields.io/github/release/homeport/yft.svg)](https://github.com/homeport/yft/releases/latest)

## Introducing YAML File Tool

The YAML File tool was originally written to serve a CLI for the [`ytbx`](https://github.com/gonvenience/ytbx) Go library. It serves as a convenience layer for the following use cases:

- Support for handling Dot-Style (`some.key.in.structure`) and Go-Patch paths (`/some/key/in/name=structure`)
- Retrieving a value from a YAML file at a specific location, described by a path
- Restructure the order of keys in a mapping to be more human friendly and readable

### Notable Use Cases

- Retrieve a value from the YAML file:

  ```sh
  $ yft get foobar.yml /list/name=one/somekey
  foobar
  ```

- Restructure the order of keys in a mapping context to be more readable:

  ```sh
  $ yft get restructure.yml /
  releases:
  - sha1: 5ab3b7e685ca18a47d0b4a16d0e3b60832b0a393
    name: binary-buildpack
    version: 1.0.32
    url: https://bosh.io/d/github.com/cloudfoundry/binary-buildpack-release?v=1.0.32

  $ yft restructure restructure.yml
  releases:
  - name: binary-buildpack
    url: https://bosh.io/d/github.com/cloudfoundry/binary-buildpack-release?v=1.0.32
    version: 1.0.32
    sha1: 5ab3b7e685ca18a47d0b4a16d0e3b60832b0a393
  ```

- List all paths in a YAML file:

  ```sh
  $ yft paths foobar.yml
  /list/name=one/somekey
  /list/name=two/somekey
  /map/foobar/key
  ```

- Create a list of paths that are common in two YAML files:

  ```sh
  $ yft compare foo.yml bar.yml
  /list/name=one/somekey
  ```

## How do I get started

There are different ways to get `yft`. You are free to pick the one that makes the most sense for your use case.

- On macOS systems, a Homebrew tap is available to install `yft`:

  ```sh
  brew install homeport/tap/yft
  ```

- Use a convenience script to download the latest release to install it in a suitable location on your local machine:

  ```sh
  curl -sL https://git.io/JvUjf | bash
  ```

## Contributing

We are happy to have other people contributing to the project. If you decide to do that, here's how to:

- get Go (`yft` requires Go version 1.17 or greater)
- fork the project
- create a new branch
- make your changes
- open a PR.

Git commit messages should be meaningful and follow the rules nicely written down by [Chris Beams](https://chris.beams.io/posts/git-commit/):
> The seven rules of a great Git commit message
>
> 1. Separate subject from body with a blank line
> 1. Limit the subject line to 50 characters
> 1. Capitalize the subject line
> 1. Do not end the subject line with a period
> 1. Use the imperative mood in the subject line
> 1. Wrap the body at 72 characters
> 1. Use the body to explain what and why vs. how

### Running test cases and binaries generation

There are multiple make targets, but running `all` does everything you want in one call.

```sh
make all
```

### Test it with Linux on your macOS system

Best way is to use Docker to spin up a container:

```sh
docker run \
  --interactive \
  --tty \
  --rm \
  --volume $GOPATH/src/github.com/homeport/yft:/go/src/github.com/homeport/yft \
  --workdir /go/src/github.com/homeport/yft \
  golang:1.17 /bin/bash
```

## License

Licensed under [MIT License](https://github.com/homeport/yft/blob/main/LICENSE)
