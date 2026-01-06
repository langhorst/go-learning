# 8. Stateful HTTP

## 8.1. Choosing a session manager

- There are many [security considerations](https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html) when working with sessions
- Use a third-party package: 
  - `gorilla/sessions`
    - Most established and well-known
    - Simple, easy-to-use API
    - Lets you store session-data client-side (in signed and encrypted cookies) or server-side (in a database)
    - _Importantly, does not_ provide a mechanism to renew session IDs (which is necessary to reduce risks associated with session fixation attacks if you're using one of the server-side session stores)
  - `alexedwards/scs`
    - Lets you store session data server-side only
    - Supports automatic loading and saving of session data via middleware
    - Has a nice interface for type-safe manipulation of data
    - _Does_ allow renewal of session IDs
- Use `gorilla/sessions` if you're storing session IDs client-side, and `alexedwards/scs` if you're storing them server-side


## 8.2. Seting up the session manager

- `alexedwards/scs` if you're going to use it in production:
  - [Documentation](https://github.com/alexedwards/scs)
  - [API Reference](https://pkg.go.dev/github.com/alexedwards/scs/v2)
- `sessions` table:

```sql
USE snippetbox;
CREATE TABLE sessions (
    token CHAR(43) PRIMARY KEY,
    data BLOB NOT NULL,
    expiry TIMESTAMP(6) NOT NULL
);
CREATE INDEX sessions_expiry_idx ON sessions (expiry);
```
  - `token` contains a unique, randomly-generated identifier for each session
  - `data` field will contain the actual session data that you want to share between HTTP requests
  - `expiry` contains an expiry time for the session
- Note that `scs.New()` function returns a pointer to a `SessionManager` struct which holds the configuration settings for your sessions
- Without using alice
  - If not using `justinas/alice` to help manage your middleware chains, you'd need to use the `http.HandlerFunc()` adapter to convert your handler functions like `app.home` to an `http.Handler`, then wrap that with session middleware instead ... ex:
  
```go
mux := http.NewServeMux()
mux.Handle("GET /{$}", app.sessionManager.LoadAndSave(http.HandlerFunc(app.home)))
mux.Handle("GET /snippet/view/:id", app.sessionManager.LoadAndSave(http.HandlerFunc(app.snippetView)))
// ... etc
```


## 8.3. Working with session data

- Behind the scenes of session management
  - _session cookie_ contains a _session token_, also known as the _session ID_
