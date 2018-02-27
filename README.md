
# Go Lang Reference Project

A go-to example Go-lang project referencing some main points of the language, all in one place for convenience.

This reference project is a personal one, whilst I've tried to ensure that the practices within are standard language practices, I am not a Go expert.

Please view both the details below and the actual source code itself as needed. Enjoy!

## Prerequisites

The reference assumes:

- `go` is installed correctly
- `$GOPATH` is correctly setup and configured (with `src`, `bin` and `pkg` within)
- A basic working knowledge of `*nix` shell commands

## Technical

### Main Go commands

- go build
- go test
- go install
- go run

### Code Standards

- gofmt
- golint

### Software Packages

- github.com

### Example Project Workspace/Layout

> main.go

### Example Lib

- structure
    - folder of package name
    - same package name for multiple files
    - file names don't actually matter but convention is usually the package name for the main file

- interface
- struct
- methods
- constructor func
- helper funcs?

#### Importing Libs

- $GOPATH
- absolute import paths

### Testing

#### Go Test

> main_test.go

#### GoConvey

> lib_test.go

### Debugging

#### Print Statements

```go
fmt.Println(some_var)
```

#### Breakpoints

- command line debugger lib?
- atom
- gogland

## Thanks

Thanks for looking, if you have any suggestions on improving the reference or new additions then feel free to raise an issue. Also, if you'd like to have your own reference project then feel free to use this as a starter and fork it.
