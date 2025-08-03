# 1. Tutorial

## 1.1 Hello, World

Package `main` is special. It defines a standalone executable program, not a library. Within his package, the *function* `main* is also special -- it's where the execution of the program begins.

Whatever main does is what the program does.

`import` declarations must follow the `package` declaration. After that, the program consists of the declarations of *functions*, *variables*, *constants*, and *types*.

A function declaration consists of the keyword `func`, the name of the function, a parameter list (empty for `main`), a result list (also empty here), and the body of the function -- the statements thta define what it does -- enclosed in braces.

Go does not require semicolons. It manages them itself.

Code formatting is handled by the `gofmt` tool. There is also a separate tool available at `golang.org/x/tools/cmd/goimports` that handles the insertion and removal of import declarations as needed.

## 1.2 Command-Line Arguments

The variable `os.Args` is a *slice* of strings.

Slices are a dynamically sized sequence `s` of array elements where individual elements can be accessed as `s[i]` and a continguious subsequence as `s[m:n]`.

The number of elements of a slice can be obtained with `len(s)`.

The slice `s[m:n]` where `0 <= m <= n <= len(s)` contains `n-m` elements.

Comments begin with `//` and end at the end of the line. 

The `var` declaration declares two `string` variables `s` and `sep`.

`+` concatenates strings. `+=` is an assigment operator.

## 1.3 Finding Duplicate Lines

A *map* holds a set of key/value pairs and provides constant-time operations to store, retrieve, or test for an item in the set. The key may be of any type whose values can be compared with `==`, strings being the most common example; the value may be of any type at all.

```go
counts[input.Text()]++
```

is the equivalent of:

```go
line := input.Text()
counts[line] = counts[line] + 1
```

You don't have to worry about the initialization of map keys. Just start using them. Also, map keys are not ordered.

## 1.4 Animated GIFs

A `const` declaration, like `var`, can be made at the package level or within functions.

The expressions `[]color.Color{...}` and `gif.GIF{..}` are *composite literals*, a compact notation for instantiating any of Go's composite types from a sequence of element values.

A *struct* type is a group of values called *fields*, often of different types, that are collected together in a single object that can be treated as a unit.
