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
## 9.6. Connection timeouts
