# 6. Middleware

- Shared functionality for many (or all) HTTP requests
  - Log every Request
  - Compress every response
  - Check cache before passing the request to handlers

## 6.1. How middleware works

> You can think of a Go web application as a chain of `ServeHTTP()` methods being called one after another.

- The pattern:

```go
func myMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// TODO: Execute our middleware logic here...
		next.ServeHTTP(w, r)
	}
	
	return http.HandlerFunc(fn)
}
```

  - The `myMiddleware()` function is essentially a wrapper around the `next` handler, which we pass to it as a parameter
  - It establishes a function `fn`, which _closes_ over the `next` handler to form a closure.
  - When `fn` is run, it executes our middleware logic and then transfers control to the `next` handler by calling its `ServeHTTP()` method
  - Regardlress of what you do with a closure, it will always be able to access the variables that are local to the scope it was created in -- which in this case means that `fn` will always have access to the `next` variable
  - Final line of code, convert this closure to an `http.Handler` and return it using the `http.HandlerFunc()` adapter
- Simplifying the middleware
  - A tweak to this pattern, using an _anonymous function_:
  
```go
func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Execute our middleware logic here...
		next.ServeHTTP(w, r)
	}
}
```

  - This is a much more common pattern in the wild
- Positioning the middleware
  - Before servemux: it acts on every request that your application receives
  - After the servemux: only executes for a specific route
  

## 6.2. Setting common headers
## 6.3. Request logging
## 6.4. Panic recovery
## 6.5. Cmoposable middleware chains
