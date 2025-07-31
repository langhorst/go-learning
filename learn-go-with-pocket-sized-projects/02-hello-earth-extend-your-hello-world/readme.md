# 2. Hello, earth! Extend your hello, world

## 2.1 Any travel begins at home

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
