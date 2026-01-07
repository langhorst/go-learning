# 9. Server and security improvements

## 9.1. The http.Server struct

- `http.ListenAndServe()` is a shortcut function
- Creating out own `http.Server` struct allows us to customize the behavior of our server


## 9.2. The server error log

- Go's `http.Server` by default may write to the standard logger, which means that the log entries will be written to the standard error stream
- Solution is to convert our structured logger _handler_ into a `*log.Logger` which writes log entries at a specific fixed level, and then register that with the `http.Server`
  - `slog.NewLogLogger()`


## 9.3. Generating a self-signed TLS certificate

- HTTPS is essentially HTTP transmitted over a TLS (_Transport Layer Security_) connection
- `go run /opt/homebrew/Cellar/go/1.25.4/libexec/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost`
  - Generates a 2048-bit RSA key pair, which is a cryptographically secure public and private key
  - Stores the private key in a `key.pem` file, and generates a self-signed TLS certificate for the host `localhost` containing the public key, which it stores in a `cert.pem` file
  - Both private key and certificate are encoded in PEM format, which is the standard format used by most TLS implementations
- The mkcert tool
  - An alternative to `generate_cert.go`
  - Has the advantage that the generated certificates are _locally trusted_ meaning that you can use them without getting browser warnings


## 9.4. Running an HTTPS server

- Swap `srv.ListenAndServe()` with `srv.ListenAndServeTLS()`
- HTTP requests
  - Our HTTPS server _only supports HTTPS_
- HTTP/2 connections
  - A big plus of using HTTPS is that Go will automatically upgrade the connection to use HTTP/2 if the client supports it
  - This means faster page loads
- Certificate permissions
  - Generally a good idea to keep the permissions of your private keys as tight as possible and allow them to be read only by the owner or a specific group
- Source control
  - Add a rule to ignore the contents of `tls`


## 9.5. Configuring HTTPS settings

- Restrict the _elliptic curves_ that can potentially be used during the TLS handshake
- As of Go 1.25 only `tls.CurveP256` and `tls.X25519` have assembly implementations
  - The others are very CPU intensive
  - Restricting the others helps ensure our server will remain performant under heavy loads
- TLS versions
  - By default Go's HTTPS server is configured to support TLS 1.2 and 1.3
  - Change the minimum and maximum TLS versions using:
    - `tls.Config.MinVersion`
    - `tls.Config.MaxVersion`
    - TLS version constants in `crypto/tls`
    
```go
tlsConfig := &tlsConfig{
	MinVersion: tls.VersionTLS10,
	MaxVersion: tls.VersionTLS12,
}
```

- Cipher suites
  - Defined in the `crypto/tls` package constants
  - Balancing security and backwards-compatability
  - Refer to Mozilla's [recommended configurations](https://wiki.mozilla.org/Security/Server_Side_TLS)


## 9.6. Connection timeouts

- IdleTimeout
  - Go enables keep-alives on all accepted connections by default
  - This setting can help reduce latency (especially for HTTPS connections) because a client can reuse the same connection for multiple requests without having to repat the TLS handshake
  - By default, keep-alive connections will be automatically closed after a couple of minutes (exact time depends on operating system)
  - There is no way to _increase_ this default (unless you write your own `net.Listener`), but you can _reduce_ it with `IdleTimeout`
- ReadTimeout
  - Setting `ReadTimeout` to 5 seconds means that if the request headers or body are still being read 5 seconds after the request is first accepted, then Go will close the underlying connection
  - Setting a short `ReadTimeout` period helps to mitigate the risk from slow-client attacks -- such as Slowloris -- which could otherwise keep a connection open indefinitely by sending partial, incomplete, HTTP(S) requests
  - If you set `ReadTimeout` but not `IdleTimeout`, then `IdleTimeout` will default to using the same setting as `ReadTimeout`
    - Avoid ambiguity and explicitly set both
- WriteTimeout
  - The `WriteTimeout` setting will close the underlying connection if our server attempts to write to the connection after a given period
  - Behaves slightly differently depending on the protocol being used:
    - HTTP: if some data is written to the connection more than the timeout period after the _read of the request header_ finished, Go will close the underlying connection instead of writing the data
    - HTTPS: if some data is written to the connection more than the timeout period after the request is _first accepted_, Go wil close the underlying connection instead of writing the data
      - This means that if you're using HTTPS it's sensible to set `WriteTimeout` to a value greater than `ReadTimeout`
    - Generally not to prevent long-running handlers, but to prevent the data that the handler returns from taking too long to write
  - The ReadHeaderTimeout setting
    - Applies to the read of the HTTP(S) headers only
  - The MaxHeaderBytes setting
    - Used to control the maximum number of bytes the server will read when parsing request headers
