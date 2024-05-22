# Cgo impress terminal

This is a part of cross-platform GUI Library for Go. See https://github.com/codeation/impress

The Cgo terminal is a Go version of the impress terminal with minimal C code (GTK+ 3 library, etc.).

Reasons to have a version of Cgo besides the C version:

- Clean logic of using GTK+ 3.
- A place to inject high-level code on the terminal side for debugging or benchmarking.

GTK+ 3 binding package [![PkgGoDev](https://pkg.go.dev/badge/github.com/codeation/itlog/gtk)](https://pkg.go.dev/github.com/codeation/itlog/gtk)

### To run this example on Debian/ Ubuntu:

0. Install `gcc`, `make`, `pkg-config` if you don't have them installed.

1. Install GTK+ 3 libraries if you don't have them installed:

```
sudo apt-get install libgtk-3-dev
```

2. Build impress terminal from source:

```
git clone https://github.com/codeation/itlog.git
cd it
go build -o it-log ./cmd
cd ..
```

3. Then run example:

```
git clone https://github.com/codeation/impress.git
cd impress
IMPRESS_TERMINAL_PATH=../itlog/it-log go run ./examples/simple/
```

Steps 0-2 are needed to build a Cgo version of impress terminal.

## Project State

### Notes

- This project is still in the early stages of development and is not yet in a usable state.
- The project tested on Debian 12.5
