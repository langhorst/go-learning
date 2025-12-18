# Chapter 1: Setting Up Your Go Environment


```
go version go1.24.2 darwin/arm64
```

## Your First Go Program

### Making a Go Module

```bash
$ go mod init hello_world
```

- A module is not just source code.
    - Also an exact specification of the dependencies of the code within the module.
    - Every module has a `go.mod` file in its root directory.
    - Running `go mod init` creates this file for you.
- The `go.mod` file declares:
    - The name of the module,
    - The minimum supported version of Go for the module,
    - And any other modules that your module depends on.

### go fmt

`go fmt` is a command-line tool that formats Go source code according to the official Go style guide. It can be used to automatically format your code, making it easier to read and maintain.

### go vet

`go vet` is a command-line tool that checks Go source code for suspicious constructs, such as Printf calls whose arguments do not align with the format string. It can be used to catch potential errors and improve code quality.
