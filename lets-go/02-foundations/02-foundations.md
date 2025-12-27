# 2. Foundations

## 2.1. Project setup and creating a module

```bash
go mod init snippetbox.justinlanghorst.com
```
```
```

- If you're creating a project which can be downloaded and used by other people and programs, it's good practice for your module path to equal the location that the code can be downloaded from: 
  - URL: `https://github.com/foo/bar`
  - Module path: `github.com/foo/bar`

## 2.2. Web Application Basics

- Three absolute essentials:
  - Handler
    - A bit like controllers in MVC patterns
    - Responsible for executing your application logic and for writing HTTP response headers and bodies
  - Router (or servemux in Go terminology)
    - Stores a mapping between the URL routing patterns for your application and the corresponding handlers
    - Usually you have one servemux for your application containing all your routes
  - Web Server
    - In Go, you can establish a web server and listen for incoming requests _as part of your application itself_
- home handler function
  - Just a regular Go function with two parameters
  - `http.ResponseWriter` parameter provides methods for assembling an HTTP response and sending it to the user
  - `*http.Request` parameter is a pointer to a struct which holds information about the current request (like the HTTP method and the URL being requested)
- Go's servemux treats the route pattern `"/"` like a catch-all
  - Which means at the moment _all_ HTTP requests to our server will be handled by the `home` function, regardless of their URL path
- TCP network address passed-in to `http.ListenAndServe` should be in the format of `"host:port"`

```bash
$ go run .
$ go run main.go
$ go run snippetbox.justinlanghorst.com
```

## 2.3. Routing Requests

| **Route Pattern** | **Handler** | **Action** |
| --- | --- | --- |
| `/` | `home` | Display the home page |
| `/snippet/view` | `snippetView` | Display a specific snippet |
| `/snippet/create` | `snippetCreate` | Display a form for creating a new snippet |

- Go's servemux has different matching rules depending on whether or not a route pattern ends with a trailing slash
  - Without slash: it will only be matched when the request URL path exactly matches the pattern in full
  - With slash: `/` or `/static/` is a _subtree path pattern_ and matched whenever the _start_ of a request URL path matches the subtree path 
- To prevent subtree path patterns from acting like they have a wildcard at the end, append special character sequence `{$}` to the end of the pattern
  - `"/{$}"` or `"/static/{$}"`
  - `"/{$}"` effectively means _match a single slash, followed by nothing else_
- Request URL paths are automatically sanitized
  - If request path contains any `.` or `..` elements or repeated slashes, the user will automatically be redirected to an equivalent clean URL
  - `/foo//bar/../baz/./` will automatically be sent a `301 Permanent Redirect` to `/foo/baz/` instead
- If a subtree path has been registered and a request is received for that subtree path _without_ a trailing slash, then the user will automatically be sent a `301 Permanent Redirect` to the subtree path with the slash added
  - Example: registered the subtree path `/foo/`, then any request to `/foo` will be redirected to `/foo/`
- Host name matching:
  - `mux.HandleFunc("foo.example.org/", fooHandler)`
  - `mux.HandleFunc("bar.example.org/", barHandler)`
  - `mux.HandleFunc("/baz", bazHandler)`
  - Any host specific patterns will be checked first and if there is a match the request will be dispatched to the corresponding handler
  - Only when there isn't a host specific match found will the non-host-specific patterns also be checked
- Default servemux
  - `http.Handle()` and `http.HandleFunc()` functions allow you to register routes _without_ explicitly declaring a servemux
  - These functions register their routes with something called the default servemux
  - Not recommended because:
    - It feels more "magic" than declaring and using your own locally-scoped servemux
    - `http:.DefaultServeMux` is a global variable in the standard library, so _any_ God code in your project can access it and potentially register a route
    - Any _third-party packages that your application imports_ can register routes with `http.DefaultServeMux` too

## 2.4. Wildcard route patterns

- Possible to define route patterns that contain _wildcard segments_
- Create more flexible routing rules
- Pass variables to your Go application via a request URL
- Wildcard segments in a route pattern are denoted by a wildcard _identifier_ inside `{}` brackets:
  - `mux.HandleFunc("/products/{category}/item/{itemID}", exampleHandler)`
    - Two wildcard segments: `category` and `itemID`
  - Each path segment can only contain one wildcard and the wildcard needs to fill the _whole_ path segment:
    - Invalid: `"/products/c_{category}"` or `"/date/{y}-{m}-{d}"`
  - Retrieve the corresponding value for a wildcard segment using its identifier and the `r.PathValue()` method:
    - `category := r.PathValue("category")`
    - `itemID := r.PathValue("itemID")`

| **Route pattern** | **Handler** | **Action** |
| --- | --- | --- |
| `/{$}` | `home` | Display the home page |
| `/snippet/view/{id}` | `snippetView` | Display a specific snippet |
| `/snippet/create` | `snippetCreate` | Display a form for creating a new snippet |

- Overlapping wildcard patterns:
  - _The most specific route pattern wins_
    - `/post/edit` _only_ matches requests with the exact path
    - whereas `/post/{id}` matches requests with the path `/post/edit`, `/post/123`, `/post/abc`, etc.
  - You can register patterns in any order and it won't affect how the servemux behaves
  - Potential edge case with two overlapping route patterns but neither one is obviously more specific than the other
    - `/post/new/{id}` and `/post/{author}/latest` overlap because they both match the request path `/post/new/latest`
      - Go's servemux considers the patterns to be in _conflict_ and will panic at runtime when initializing the routes
  - Best practice is to keep overlaps to a minimum or avoid them completely
- Subtree path patterns with wildcards
  - If you don't want this behavior, stick a `{$}` at the end
- Remainder wildcards
  - If a route pattern ends with a wildcard, and this final wildcard identifier ends in `...`, then the wildcard will match any and all remaining segments of a request path
    - Example: `"/post/{path...}"` matches:
      - `/post/a`
      - `/post/a/b`
      - `/post/a/b/c` and so on

## 2.5. Method-based routing

- You can restrict routes to match specific HTTP methods
  - In fact, this SHOULD be done
  - Sets the foundation for a secure web application
- We want to make sure we do two things:
  - Routes which only return data, without changing anything in the application, only match requests with the HTTP method GET
  - Routes that modify something in the application (or in other words, _change the state of the server_) only match requests with the HTTP method POST
- The HTTP methods in route patterns are case sensitive and should always be written in uppercase, followed by at least one whitespace character (both spaces and tabs are fine)
  - You can only include one HTTP method in each route pattern
- GET will match both GET and HEAD requests while all other methods (POST, PUT, DELETE) require an exact match
- Adding a POST-only route and handler

| **Route pattern** | **Handler** | **Action** |
| --- | --- | --- |
| `GET /{$}` | `home` | Display the home page |
| `GET /snippet/view/{id}` | `snippetView` | Display a specific snippet |
| `/snippet/create` | `snippetCreate` | Display a form for creating a new snippet |
| `POST /snippet/create` | `snippetCreatePost` | Save a new snippet |

- The _most specific pattern wins_ rule also applies if you have route patterns that overlap because of an HTTP method
- Route patterns that don't include a method will match incoming HTTP requests with _any method_

## 2.6. Customizing responses

- It's only possible to call `w.WriteHeader()` once per response, and after the status code has been written it can't be changed
- If you don't call `w.WriteHeader()` explicitly, then the first call to `w.Write()` will automatically send a 200 status code to the user
  - If you want to send a non-200 status code, you must call `w.WriteHeader()` _before_ any call to `w.Write()`
- Status code constants
  - `net/http` provides constants for HTTP status codes
  - Good practice to prevent mistakes due to typos
  - Helps make your code clearer and self-documenting
- Customizing headers
  - By changing the _response header map_
  - You must make sure that your response header map contains all the headers you want _before_ you call `w.WriteHeader()` or `w.Write()`
  - Any changes you make to the response header map after calling `w.WriteHeader()` or `w.Write()` will have no effect on the headers that the user receives
- Writing response bodies
  - Common to pass `http.ResponseWriter` value to _another function_ that writes the response for you
  - _because the_ `http.ResponseWriter` _value in your handlers has a_ `Write()` _method, it satisfies the_ `io.Writer` _interface_
  - This means you can use standard library functions `io.WriteString()` and the `fmt.Fprint*()` family (all of which accept an `io.Writer` parameter) to write plain- text response bodies too

```go
// Instead of this...
w.Write([]byte("Hello world"))

// You can do this...
io.WriteString(w, "Hello world")
fmt.Fprint(w, "Hello world")
```

- Additional information
  - In order to automatically set the Content-Type header, Go _content sniffs_ the response body with the `http.DetectContentType()` function
    - If this function can't guess the content type, Go will fallback to setting the header `Content-Type: application/octet-stream` instead
  - `http.DetectContentType()` can't distinguish JSON from plain-text
    - By default, they will be sent with a `Content-Type: text/plain; charset=utf-8` header
    - Prevent this from happening by setting the correct header manually:

```go
w.Header().Set("Content-Type", "application/json")
w.Write([]byte(`{"name":"Alex"}`))
```

  - Manipulating the header map
    - `Add()`, `Set()`, `Del()`, `Get()` and `Values()`

```go
// Set a new cache-control header. If an existing "Cache-Control" header exists
// it will be overwritten.
w.Header().Set("Cache-Control", "public, max-age=31536000")

// In contrast, the Add() method appends a new "Cache-Control" header and can
// be called multiple times.
w.Header().Add("Cache-Control", "public")
w.Header().Add("Cache-Control", "max-age=31536000")

// Delete all values for the "Cache-Control" header.
w.Header().Del("Cache-Control")

// Retrieve the first value for the "Cache-Control" header.
w.Header().Get("Cache-Control")

// Retrieve a slice of all values for the "Cache-Control" header.
w.Header().Values("Cache-Control")
```

  - Header canonicalization
    - When using these methods on the header map, the header name will always be canonicalized using the `textproto.CanonicalMIMEHeaderKey()` function
      - Converts the first letter and any letter following a hyphen to upper case, and the rest of the letters to lowercase
      - Header name is _case-insensitive_

## 2.7. Project structure and organization

- _Don't over-complicate things_
- Structure:
  - `cmd`: contains _application-specific_ code for the executable applications in the project
  - `internal`: contains the ancillary _non-application-specific_ code used in the project
    - Potentially reusable code like validation helpers and the SQL database models for the project
  - `ui`: contains the _user-interface assets_ used by the web application
    - `ui/html`: contains HTML templates
    - `ui/static`: contains static files (like CSS and images)
- Why?
  - Clean separation between Go and non-Go assets_
  - Scales nicely if you want to add another executable application to your project
    - Like a future CLI to automate some administrative tasks in the future
      - `cmd/cli` for example
- Additional information
  - The internal directory
    - `internal` carries a special meaning and behavior in Go: any packages which live under this directory can only be imported by code _within the parent directory of_ `internal`
      - Or in other words, packages under `internal` _cannot be imported by code outside of our project_

## 2.8. HTML templating and inheritance

- The file path passed to `template.ParseFiles()` must either be relative to your current working directory, or an absolute path
- If either `template.ParseFiles()` or `ts.Execute()` returns an error, we log the detailed error message and then use the `http.Error()` function to send a response to the user
  - `http.Error()` is a lightweight helper function which sends a plain-text error message and a specific HTTP status code to the user
- Template composition
  - Prevent duplication
