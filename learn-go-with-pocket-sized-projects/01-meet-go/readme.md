# 1. Meet Go

## What is Go?

- Go addresses:
    - Slow program construction
    - Out-of-control dependency management
    - Complex code
    - Difficult cross-language construction.
- Targets modern engineering:
    - Removing the constraint of dealing with memory
    - Making it simple to run parallel pieces of code
- Toolchain covers:
    - Compilation and construction
    - Code formatting
    - Package dependency management
    - Static code inspection
    - Testing
    - Document generation and viewing
    - Performance analysis
    - Language servers
    - Runtime program tracking
    - (more)
- Built for:
    - Concurrency
    - Networked servers

> Go started in September 20007 when Robert Griesemer, Ken Thompson, and I began discussing a new language to address the engineering challenges we and our colleagues at Google were facing in our daily work.
-- Rob Pike, coauthor of Go

- Design choices driven primarily by *simplicity*
    - Only 25 reserved keywords (Nov 2024)

## 1.2 Why you should learn Go

- Natural fit for:
    - Backend services
    - APIs
    - Modern cloud computing needs

### 1.2.1 How and where Go can help you

- Designed for maintainability and readability
- Optimal for backend software development
- Has great integration with modern cloud technologies
- A reliable, secure language with a fast build time
- To make stacks small, Go uses resizable, bounded stacks
- Newly minted goroutines:
    - Given a few kilobytes
    - Can grow and shrink
- It's practical to create hundreds of thousands of goroutines in the same address space
- Supports most features of object-oriented programming via composition and implicit interfaces
    - No inheritance system
- Generics in 1.18 (finally)
    - Make it possible to write reusable, type-safe functions
    - Significantly reduce boilerplate code

### 1.2.2 Where Go cannot help you

- Relies on a garbage collector to release memory
    - If you require full control over memory, use C/C++ or Rust
- Go can wrap libraries written in C/C++ with cgo
    - Translation layer created to ease transition between the two languages
- TinyGo for embedded use

## 1.3 Why pocket-sized projects?

- 1897. John Dewey. *Doing* is the best way to learn.

### 1.3.1 What you'll know after reading the book (and writing the code)

- Good and clear examples for writing industry-level Go code with recommendations toward real-world dev
- Functions reusable in a professional codebase

#### Grammar ad Syntax

- Interfaces are *implicit*
    - Therefore you can unknowingly implement interfaces
    - Implementing methods is all you need to do
- Won't dwell on goroutines
- Errors as values

#### Testing Your Code

- Tests are indispensable
- Unit tests are included everywhere
- Fuzzing!

#### Clean Code Best Practices

> Any code of your own that you haven't looked at for six or more months might as well have been written by someone else.
-- Eagleson's Law

- Go is great at a domain-driven design, so we organize our code accordingly

#### Architectural Decisions

- Because Go is great for writing services deployed in cloud environments, two projects:
    - One serves HTML over HTTP
    - The other uses Protobuf over gRPC

#### Your Go Toolbox

- Go isn't just a compiler; it's a complex suite of tools all integrated for:
    - Linting
    - Formatting
    - Testing
    - Benchmarking
    - Building
    - Executing
- Ventures into the world of other architectures, including microcontrollers
- Covers shipping a Go program as Web Assembly in a web page

#### Side Quests

- Designed to practice and think more deeply about Go


