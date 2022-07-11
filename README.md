# Go-Mosh: Golang Wrapper for Mosh (Mobile Shell)

> Mosh is a remote terminal application that supports intermittent connectivity, allows roaming, and provides
> speculative local echo and line editing of user keystrokes.

Go-Mosh is a SWIG based wrapper so that (parts of) Mosh can be used as a library from Golang. Mosh is licensed under the
terms of the GPLv3, as is this wrapper.

This SWIG specification wraps only a small subset of Mosh, specifically the following C++ classes:

- `Overlay::PredictionEngine`
- `Parser::Action`
- `Parser::Resize`
- `Parser::UserByte`
- `Terminal::Complete`
- `Terminal::Display`
- `Terminal::Framebuffer`

These are essentially intended to make Mosh's speculative and predictive terminal emulation capabilities available as a
library.

The contents of the `internal` directory are the directly exposed C++ classes. The contents of the `pkg` directory are
wrapped versions of those C++ classes that are better interoperable with idiomatic Go code (i.e. they are automatically
garbage collected, and where appropriate accept/return Go objects e.g. `time.Time` and `time.Duration`).

## Required Runtime Libraries

- `ncurses`
- `protobuf`
- `crypto` (Linux)

## Instructions

- Install SWIG (tested with versions 3.0.9, 4.0.0, 4.0.2)
- Install required runtime libraries
- Run `git submodule update --init` to retrieve a copy of the Mosh repository
- Build Mosh by following the instructions within the Mosh repository
- Use go-mosh library and build downstream project normally with `go build`

Build is confirmed functional on macOS (darwin) and Linux platforms when using Clang toolchain.

Note: When using go-mosh as a Go module, follow these installation instructions within the instance of go-mosh loaded
into the module managed $GOPATH.
