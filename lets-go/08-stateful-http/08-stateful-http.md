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
## 8.3. Working with session data
