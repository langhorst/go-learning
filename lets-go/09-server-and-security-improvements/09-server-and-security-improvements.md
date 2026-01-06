# 9. Server and security improvements

## 9.1. The http.Server struct

- `http.ListenAndServe()` is a shortcut function
- Creating out own `http.Server` struct allows us to customize the behavior of our server


## 9.2. The server error log

- Go's `http.Server` by default may write to the standard logger, which means that the log entries will be written to the standard error stream
- Solution is to convert our structured logger _handler_ into a `*log.Logger` which writes log entries at a specific fixed level, and then register that with the `http.Server`
  - `slog.NewLogLogger()`


## 9.3. Generating a self-signed TLS certificate
## 9.4. Running an HTTPS server
## 9.5. Configuring HTTPS settings
## 9.6. Connection timeouts
