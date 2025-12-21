# 2. Foundations

## Project setup and creating a module

```bash
go mod init snippetbox.justinlanghorst.com
```
```
```

- If you're creating a project which can be downloaded and used by other people and programs, it's good practice for your module path to equal the location that the code can be downloaded from: 
  - URL: `https://github.com/foo/bar`
  - Module path: `github.com/foo/bar`

## Web Application Basics

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

## Routing Requests

| **Route Pattern** | **Handler** | **Action** |
| --- | --- | --- |
| `/` | `home` | Display the home page |
| `/snippet/view` | `snippetView` | Display a specific snippet |
| `/snippet/create` | `snippetCreate` | Display a form for creating a new snippet |
