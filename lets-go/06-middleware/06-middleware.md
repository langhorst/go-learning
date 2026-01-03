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

- Headers:
  - `Content-Security-Policy` (CSP) headers are used to restrict where the resources for your web page (e.g. JavaScript, images, fonts, etc) can be loaded from
    - Setting a strict CSP policy helps prevent a variety of cross-site scripting, clickjacking, and other code-injection attacks
    - Primer: https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP
  - `Referrer-Policy` is used to control what information is included in a `Referer` header when a user navigates away from your web page
    - `origin-when-cross-origin` means that the full URL will be included for [same-origin requests](https://developer.mozilla.org/en-US/docs/Web/Security/Same-origin_policy), but for all other requests information like the URL path and any query string values will be stripped out
  - `X-Content-Type-Options: nosniff` instructs browsers to _not_ MIME-type sniff the content-type of the response, which in turn helps to prevent [content-sniffing attacks](https://security.stackexchange.com/questions/7506/using-file-extension-and-mime-type-as-output-by-file-i-b-combination-to-dete/7531#7531)
  - `X-Frame-Options: deny` is used to help prevent [clickjacking](https://developer.mozilla.org/en-US/docs/Web/Security/Types_of_attacks#click-jacking) attacks in older browsers that don't support CSP headers
  - `X-XSS-Protection: 0` _disables_ the browser's built-in XSS (cross-site-scripting) filtering
    - Previously it was good practice to set this header to `X-XSS-Protection: 1; mode=block` but when you're using CSP headers like we are the recommendation is to disable this
- Flow of control
  - When the last handler in the chain returns, control is passed back up the chain in the reverse direction
  - In any middleware handler
    - Code which comes before `next.ServeHTTP()` will be executed on the way down the chain
    - Code after `next.ServeHTTP()` -- or in a deferred function -- will be executed on the way back up

```go  
func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Any code here will execute on the way down the chain.
		next.ServeHTTP(w, r)
		// Any code here will execute on the way back up the chain.
	})
}
```

- Early returns
  - If you call `return` in your middleware function _before_ you call `next.ServeHTTP()`, then the chain will stop being executed and control will flow back upstream
  - Authentication middleware is a good use case:
  
```go
func myMiddleware(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		// If the user isn't authorized, send a 403 Forbidden status and
		// return to stop executing the chain
		if !isAuthorized(r) {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		
		// Otherwise, call the next handler in the chain.
		next.ServeHTTP(w, r)
	})
}
```

- Debugging CSP issues
  - Watch your browser logs and get in the habit of checking logs early


## 6.3. Request logging

- Middleware method based on `application` _also_ has access to the handler dependencies including the structured logger


## 6.4. Panic recovery
## 6.5. Cmoposable middleware chains
