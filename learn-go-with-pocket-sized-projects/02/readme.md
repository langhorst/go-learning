# 2. Hello, earth! Extend your hello, world

## Internal and external testing

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
