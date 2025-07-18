# Chapter 1: Your First Command-Line Program in Go

- This first program will give you an idea of how to build and test a command-line application in Go.

## Building the Basic Word Counter

- A tool that counts the number of words or lines provided as input using the *standard input* (STDIN) connection.
- By default, counts the number of words, unless it receives the `-l` flag, in which case it'll count the number of lines instead.

## Compiling Your Tool for Different Platforms

```bash
$ GOOS=windows go build
````

- It's incredibly easy to build for other platforms (at least in most cases)
- The documentation for `go build` contains a list with all of the supported vlues for the `GOOS` environment variable.
  - `golang.org/src/go/build/syslist.go`

## Exercises

- Exercise adds the `-b` flag to count the number of bytes instead.
