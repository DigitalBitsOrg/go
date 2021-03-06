# Frontier
[![Build Status](https://travis-ci.org/digitalbitsorg/frontier.svg?branch=master)](https://travis-ci.org/digitalbitsorg/frontier)

Frontier is the [client facing API](/docs) server for the DigitalBits ecosystem.  It acts as the interface between digitalbits-core and applications that want to access the DigitalBits network. It allows you to submit transactions to the network, check the status of accounts, subscribe to event streams, etc. See [an overview of the DigitalBits ecosystem](https://developer.digitalbits.io/guides/get-started/) for more details.

## Downloading the server
[Prebuilt binaries](https://github.com/digitalbitsorg/go/releases) of frontier are available on the 
[releases page](https://github.com/digitalbitsorg/go/releases).

See [the old releases page](https://github.com/digitalbitsorg/frontier/releases) for prior releases

| Platform       | Binary file name                                                                         |
|----------------|------------------------------------------------------------------------------------------|
| Mac OSX 64 bit | [frontier-darwin-amd64](https://github.com/digitalbitsorg/go/releases/download/frontier-v0.12.0-testing/frontier-v0.12.0-testing-darwin-amd64.tar.gz)      |
| Linux 64 bit   | [frontier-linux-amd64](https://github.com/digitalbitsorg/go/releases/download/frontier-v0.12.0-testing/frontier-v0.12.0-testing-linux-amd64.tar.gz)       |
| Windows 64 bit | [frontier-windows-amd64.exe](https://github.com/digitalbitsorg/go/releases/download/frontier-v0.12.0-testing/frontier-v0.12.0-testing-windows-amd64.zip) |

Alternatively, you can [build](#building) the binary yourself.

## Dependencies

Frontier requires go 1.6 or higher to build. See (https://golang.org/doc/install) for installation instructions.

## Building

[glide](https://glide.sh/) is used for building frontier.

Given you have a running golang installation, you can install this with:

```bash
curl https://glide.sh/get | sh
```

Next, you must download the source for packages that frontier depends upon. From within the project directory, run:

```bash
glide install
```

Then, simply run `go install github.com/digitalbitsorg/go/services/frontier`.  After successful
completion, you should find `frontier` is present in your `$GOPATH/bin` directory.

More detailed intructions and [admin guide](internal/docs/reference/admin.md). 

## Developing Frontier

See [the development guide](internal/docs/developing.md).

## Contributing
Please see the [CONTRIBUTING.md](./CONTRIBUTING.md) for details on how to contribute to this project.
