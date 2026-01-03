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
