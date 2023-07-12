# def

> A comfy dictionary navigator for the command line. ☁️

[![Latest Release](https://img.shields.io/github/release/fawni/def.svg)](https://github.com/fawni/def/releases)
[![Build Status](https://img.shields.io/github/actions/workflow/status/fawni/def/build.yml?logo=github&branch=master)](https://github.com/fawni/def/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/fawni/def)](https://goreportcard.com/report/github.com/fawni/def)

![scrot](assets/scrot.png)

## Installation

### Binaries

Download a binary from the [releases](https://github.com/fawni/def/releases)
page.

### Build from source

Go 1.18 or higher required. ([install instructions](https://golang.org/doc/install.html))

    go install github.com/fawni/def@latest

## Usage

```
def <word>
```

`def -h` for more information.

### Flags

- `-r`, `--related`: displays synonyms and antonyms if found
- `-e`, `--examples`: displays examples if found
- `-f`, `--full`: displays synonyms, antonyms and examples if found

## License

[ISC](LICENSE)
