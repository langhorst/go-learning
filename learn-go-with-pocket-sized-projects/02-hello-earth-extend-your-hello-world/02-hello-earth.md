# 2. Hello, earth! Extend your hello, world

> **Programmatic greeting history**
>
> The `hello, world` programmatic greeting was made popular by Brian Kernighan and Dennis Ritchie's _The C Programming Language_, published in 1978. The sentence originally came from another publication, also by Brian Kernighan, "A Tutorial Introduction to the Language B," published in 1972. This was, in all honesty, actually the second example of printing characters in this publication--the first one having the pogram print `hi!` The reason was that the B language limited the number of ASCII characters in a single variable to four characters. `Hello, world`, as a result, was achieved with several calls to the printing function. The original article (https://www.scribd.com/document/494619413/btut) printed `hell`; then `o, w`; then `orld`, and finally `!`; resulting in `hello, world!` This message was inspired by a bird hatching out of its egg, as shown in a comic strip.

- Good code should both be documented and tested


## 2.1 Any travel begins at home

- Naming things:
  - Maybe the biggest challenge of writing code
  - Guidelines:
    - If the scope of the variable is limited to two to three lines, a one- or two-letter placeholder is perfectly valid
      - No random letters
      - Use something that immediately reminds you of the variable's purpose
    - Stay consistent between different functions
    - Otherwise, use a name that explicitly refers to the current entity
      - Think; `row`, `column`, `book`, `address`, `order` and so on
    - Go's convention is to use camelCase for unexposed functions, types, variables, and constants
      - For packages, try as much as you can to use a single word
    - Go's variables don't need to describe their type -- Hungarian notation isn't in use in Go
    - Variable names can't start with a number, nor can functions, types, or constants
- A capital question;
  - Any symbol starting with a capital letter is exposed to external users of the package
  - Anything else isn't accessible from outside the package


```bash
go mod init example.com/your-repository
```

or

```bash
go mod init learngo-pockets/hello
```

### The Example function

This is from Claude:

- `Example()` for package-level examples
- `ExampleFunction()` for examples of a specific function named `Function`
- `ExampleType()` for examples of a specific type named `Type`
- `ExampleType_Method()` for examples of a method on a type

I'm going to make the assumption that this is correct, especially because `ExampleMain()` flat out did not work from the book, while `Example()` did.

Either way, the fact that the testing framework supports this type of testing is great:

```go
func Example() {
	main()
	// Output:
	// Hello world
}
````

### Internal and external testing

Two approaches:

- *External testing*
  - Test from the user's point of view
  - We can only test what is exposed
  - `{packagename}_test` and in the same folder
- *Internal testing*
  - Test unexposed functions
  - Test files should should e in the same package as the source file

Using `ExampleMain()` to test the `main()` function didn't work. Maybe this is supposed to be a special case, or was a special case at some point. I'm not sure, but in `go1.24.2` this didn't work.

## Every test has four main steps

- *Preparation phase*
  - set up everything we need to run the test:
    - input values
    - expected outputs
    - environment variables
    - global variables
    - network connection
- *Execution phase*
  - call the tested function
  - usually a single line
- *Decision phase*
  - check that the output we got corresponds to the output we want
  - might include:
    - several comparisons,
    - evaluations
    - sometimes some processing
    - test failing or passing
- *Teardown phase*
  - clean back to whatever the state was prior to the test's Execution
  - extremely simple due to `defer` keyword
    - anything altered or created during preparation should be fixed or destroyed here

## 2.2 Are you a polyglot?

### Clarity through typing

```go
type language string
```

This may seem a little silly at first, but type definitions help understanding of what values to expect and makes mixing up parameters more difficult.

### Use Test<FunctionName> functions

The first tests of multiple languages used test functions for `greet(l language)` in the following format:

```go
func TestGreet_English(t *testing.T) {}
func TestGreet_French(t *testing.T) {}
```

... and so on.

## 2.3 Supporting more languages with a phrasebook

Table-driven testing utilizing a map hash table.

## 2.4 Using the flag package to read the user's language

Great, but Go has always been weird with their default package. While it's great for small utilities, I think using spf13/cobra for most things is probably the better choice.

## Summary

- `go run` allows you to quickly execute a Go program without creating a binary, making it useful during development.
- Writing tests alongside your code, rather than after it, ensures the mental model of expected inputs and outputs is fresh and accurate, leading to better test coverage.
- `go test` is used to execute tests written for your code, following Go's naming conventions for test files and functions.
- The `testing` package provides everything you need to write and run tests, including support for assertions and benchmarks.
- Table-driven tests are a Go best practice for testing functions across multiple inputs and outputs. By using `slices` of test cases, you can easily iterate through various scenarios in a structured way.
- `Example` functions are special tests used to check the output of a function by verifying what is written to the standard output. These are often used to demonstrate usage.
- Each test generally follows four phases: preparation, where you set up inputs; execution, where you run the code being tested; decision, where you compare the result with exepected output; and teardown, where you clean up resources.
- The `flag` package allows for parsing command-line arguments, making your Go applications more flexible and interactive.
- In Go, `map`s are powerful data structures for storing key-value pairs. Accessing a `map` returns both the value and a Boolean indicating if the key was found.
- Define custom types when they provide meaningful context over built-in types. For example, a `UserID` type can be more descriptive than an `int`. 
- Use `if` statements for simply binary conditions. For more complex cases, prefer `switch` statements or `map`s to handle multiple conditions more cleanly.

