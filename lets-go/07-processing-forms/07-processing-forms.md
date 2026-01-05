# 7. Processing forms

## 7.1. Setting up an HTML form

- Setting up an HTML form


## 7.2. Parsing form data

- Two distinct steps:
  1. Use `r.ParseForm()` to parse the request body
    - This checks that the request body is well-formed
    - Stores the form data in the request's `r.PostForm` map
    - If errors, returns an error
    - `r.ParseForm()` is idempotent
  2. Get the form data contained in `r.PostForm` by using the `r.PostForm.Get()` method
    - If no matching value, returns an empty string
- The PostFormValue method   
  - `r.PostFormValue()` method calls `r.ParseForm()` and fetches the appropriate field value from `r.PostForm`
  - Avoid this shortcut because it _silently ignores any errors_ returned by `r.ParseForm()`
- Multiple-value fields
  - `r.PostForm.Get()` only works for the _first_ value of the specific form field, so it does not work for things like checkboxes with multiple values
  - The solution is to work with the `r.PostForm` map directly, which is a map of `url.Values`
  
```go
for i, item := range r.PostForm["items"] {
	fmt.Fprintf(w, "%d: Item %s\n", i, item)
}
```

- Limiting form size
  - Forms submitted with `POST` have a size limit of 10MB of data by default
  - Exception is if form has the `enctype="multipart/form-data"` attribute and sends multipart data (no default limit here)
  - Change 10MB limit with `http.MaxBytesReader()`:
  
```go
// Limit the request body size to 4096 bytes
r.Body = http.MaxBytesReader(w, r.Body, 4096)

err := r.ParseForm()
if err != nil {
	http.Error(w, "Bad Request", http.StatusBadRequest)
	return
}
```
  - Trying to read beyond this limit of 4096 bytes will cause the `MaxBytesReader` to return an error
  - `MaxBytesReader` sets a flag on `http.ResponseWriter` if the limit is reached, and this instructs the server to close the underlying TCP connection
- Query string parameters
  - `GET` vs `POST` will include form data in the URL _query string parameters_
  - Ex: `/foo/bar?title=value&content=value`
  - Retrieve with `r.URL.Query().Get()` method
- The r.Form map  
  - Contains form data from any `POST` request body _and_ any query string parameters
  - Can be helpful if you want your application to be agnostic about how data values are passed to it
 
  
## 7.3. Validating form data

- `utf8.RuneCountInString()` counts the number of _Unicode code points_ rather than the number of bytes


## 7.4. Displaying errors and repopulating fields

- Restful routing
  - We are not matching Ruby on Rails or Laravel here so that things are not clunky and confusing
  

## 7.5. Creating validation helpers

- Generics
  - _parametric polymorphism_
  - Allow you to write code that works with _different concrete types_
  - From this `[]string` slice and `[]int` slice handling:
  
```go
// Count how many times the value v appears in the slice s.
func countString(v string, s []string) int {
	count := 0
	for _, vs := range s {
		if v == vs {
			count++
		}
	}
	return count
}

func countInt(v int, s []int) int {
	count := 0
	for _, vs := range s {
		if v == vs {
			count++
		}
	}
	return count
}
```

  - And with generics:
  
```go
func count[T comparable](v T, s []T) int {
	count := 0
	for _, vs := range s {
		if v == vs {
			count++
		}
	}
	return count
}
```

  - Check out the [official Go generics tutorial](https://go.dev/doc/tutorial/generics) and then watch the first 15 minutes of [this video](https://www.youtube.com/watch?v=Pa_e9EeCdy8)
  - Use judiciously, you don't _need_ generics, but consider using when:
    - You find yourself writing repeated boilerplate code for different data types
      - Examples:
        - common operations on slices, maps or channels
        - helpers for carrying out validation checks
        - test assertions on different data types
    - You find yourself reaching for the `any` (empty `interface{}`) type
      - Examples:
        - creating a data structure (like a queue, cache or linked list) which needs to operate on different types
  - You probably don't want to use generics:
    - If it makes your code harder to understand, or less clear
    - If all the types that you need to work with have a common set of methods -- in which case it's better to define and use a normal `interface` type instead
    - Just _because you can_ -- prefer non-generic code by default, and switch to a generic version later _only if it is actually needed_


## 7.6. Automatic form parsing

- Panicking vs returning errors
  - The decision to panic within `decodePostForm()` if we get a `form.InvalidDecoderError` error is not taken lightly
  - It's generally considered best practice in Go to return your errors and handle them gracefully
  - In _some special circumstances_ it can be OK to panic -- no need to be dogmatic about _not panicking_ when it makes sense to
  - Two classes of errors:
    - _operational errors_ that may occur during normal operation
      - Examples:
        - caused by a database query timeout
        - a network resource being unavailable
        - bad user input
      - These errors don't necessarily mean there is a problem with your program itself -- in fact they're often caused by things outside the control of your program
      - Good practice to return these kinds of errors and handle them gracefully
    - _programmer errors_ which should not happen during normal operation, and if they do it is probably the result of a developer mistake or a logical error in your codebase
      - These are truly exceptional errors, and using panic in these circumstances is more widely accepted
      - The Go standard library frequently does this when you make a logical error or try to use the language features in an unintended way
      - Example:
        - accessing an out-of-bounds index in a slice
        - trying to close an already-closed channel
      - Even these the recommendation is to return and gracefully handle programmer errors in most cases
        - Exception to this is when _returning the error_ adds an unacceptable amount of error handling to the rest of your codebase
  - The panic in `decodePostForm()` helper: if we get a `form.InvalidDecoderError` at runtime it's because we as the developers have tried to use something that isn't a _non-nil pointer_ as the target decode destination, and this is firmly a programmer error which we _shouldn't_ see under normal operation, and is something that should be picked up in development and tests long before deployment 
  - _Go by Example_ page on panics summarizes all of this quite nicely:

> A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn't occur during normal operation and that we aren't prepared to handle gracefully.

  - More info in [this tutorial](https://www.alexedwards.net/blog/when-is-it-ok-to-panic-in-go)
