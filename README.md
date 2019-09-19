# rostamctl [![Build Status](https://github.com/khos2ow/rostamctl/workflows/build/badge.svg)](https://github.com/khos2ow/rostamctl/actions) [![GoDoc](https://godoc.org/github.com/khos2ow/rostamctl?status.svg)](https://godoc.org/github.com/khos2ow/rostamctl) [![Go Report Card](https://goreportcard.com/badge/github.com/khos2ow/rostamctl)](https://goreportcard.com/report/github.com/khos2ow/rostamctl)

rostamctl is a tool for interacting with [RostamBot](https://rostambot.com/) APIs via a command line interface.

## Installation

The latest version can be installed using `go get`:

``` bash
GO111MODULE="on" go get github.com/khos2ow/rostamctl@v0.1.0
```

**NOTE:** please use the latest go to do this, ideally go 1.12.9 or greater.

This will put `rostamctl` in `$(go env GOPATH)/bin`. If you encounter the error `rostamctl: command not found` after installation then you may need to either add that directory to your `$PATH` as shown [here](https://golang.org/doc/code.html#GOPATH) or do a manual installation by cloning the repo and run `make build` from the repository which will put `rostamctl` in:

```bash
$(go env GOPATH)/src/github.com/khos2ow/rostamctl/bin/$(uname | tr '[:upper:]' '[:lower:]')-amd64/rostamctl
```

Stable binaries are also available on the [releases](https://github.com/khos2ow/rostamctl/releases) page. To install, download the binary for your platform from "Assets" and place this into your `$PATH`:

```bash
curl -Lo ./rostamctl https://github.com/khos2ow/rostamctl/releases/download/v0.1.0/rostamctl-$(uname)-amd64
chmod +x ./rostamctl
mv ./rostamctl /some-dir-in-your-PATH/rostamctl
```

## Code Completion

The code completion for `bash` or `zsh` can be installed using:

**Note:** Shell auto-completion is not available for Windows users.

### bash

``` bash
rostamctl completion bash > ~/.rostamctl-completion
source ~/.rostamctl-completion

# or simply the one-liner below
source <(rostamctl completion bash)
```

### zsh

``` bash
rostamctl completion zsh > /usr/local/share/zsh/site-functions/_rostamctl
autoload -U compinit && compinit
```

To make this change permenant, the above commands can be added to your `~/.profile` file.
