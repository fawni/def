# def

> A comfy dictionary navigator for the command line. ☁️

![scrot](assets/scrot.png)

<!-- [![Latest Release](https://img.shields.io/github/release/x6r/def.svg)](https://github.com/x6r/def/releases)
[![Build Status](https://img.shields.io/github/workflow/status/x6r/def/build?logo=github)](https://github.com/x6r/def/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/x6r/def)](https://goreportcard.com/report/github.com/x6r/def) -->

## Installation

### Binaries

Download a binary from the [releases](https://github.com/x6r/def/releases)
page.

### Build from source

Go 1.16 or higher required. ([install instructions](https://golang.org/doc/install.html))

    go install github.com/x6r/def@latest

## Usage

```sh
$ def <word>
```

`def -h` for more information.

### Flags

- `-r`, `--related`: displays synonyms and antonyms if found
- `-e`, `--examples`: displays examples if found
- `-f`, `--full`: displays synonyms, antonyms and examples if found

## License

[ISC](LICENSE)
