# impress terminal. Developer version

This is a part of cross-platform GUI Library for Go. See https://github.com/codeation/impress

The developer version is a Go version of [the impress terminal](https://github.com/codeation/it) with minimal C code (GTK+ 3 library, etc).

Reasons to have a developer version besides the C version:

- A reference implementation of the client side.
- A place to inject high-level code on the client side for debugging or benchmarking.
- Highlight the clean logic of using GTK+ 3.

Yet another GTK+ 3 binding package inside [![PkgGoDev](https://pkg.go.dev/badge/github.com/codeation/itlog/gtk)](https://pkg.go.dev/github.com/codeation/itlog/gtk)

### To run this example on Debian/ Ubuntu:

0. Install `gcc`, `make`, `pkg-config` if you don't have them installed.

1. Install GTK+ 3 libraries if you don't have them installed:

```
sudo apt-get install libgtk-3-dev
```

2. Build impress terminal from source:

```
git clone https://github.com/codeation/itlog.git
cd itlog
go build -o itlog github.com/codeation/itlog/cmd
cd ..
```

3. Then run example:

```
git clone https://github.com/codeation/impress.git
cd impress
IMPRESS_TERMINAL_PATH=../itlog/itlog go run github.com/codeation/impress/examples/simple/
```

Steps 0-2 are needed to build a Cgo version of impress terminal.

## Project State

### Notes

- The project is currently in its beta stage.
- The project tested on Debian 12.8.
