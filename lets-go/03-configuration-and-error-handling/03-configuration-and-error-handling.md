# 3. Configuration and error handling

## 3.1. Managing configuration settings

- Hard-coding isn't ideal, and our Snippetbox has a few things hard-coded:
  - The network address for the server to listen on (`:4000`)
  - The file path for the static files directory (`./ui/static`)
  - Separate configuration and code (important if we need different settings for development, testing and production environments, for example)
- Command-line flags
  - A common and idiomatic way to manage configuration settings in Go is to use command-line flags when starting an application:
    - `$ go run ./cmd/web -addr=":80"`
    - Easiest way to do this: `addr := flag.String("addr", ":4000", "HTTP network address")`
- Default values
  - Use whatever defaults make sense
- Type conversions
  - Go has a range of other functions for defining flags that automatically convert the command-line flag value to the appropriate type:
    - `flag.Int()`
    - `flag.Bool()`
    - `flag.Float64()`
    - `flag.Duration()`
- Automated help: `$ go run ./cmd/web/ -help`
- Environment variables
  - If you want, you _can_ store your configuration settings in environment variables and access them directory from your application by using the `os.Getenv()` function: `addr := os.Getenv("SNIPPETBOX_ADDR")`
  - Drawbacks compared to using command-line flags:
    - You can't specify a default setting
    - You don't get the `-help` functionality that you do with command-line flags
    - Return value from `os.Getenv()` is _always_ a string -- you don't get the automatic type conversions like you do with the rest of the functions
  - Get the best of both worlds by passing the environment variable as a command-line flag when starting the application:

```bash
$ export SNIPPETBOX_ADDR=":9999"
$ go run ./cmd/web -addr=$SNIPPETBOX_ADDR
```

- Boolean flags
  - For flags defined with `flag.Bool()`, omitting a value when starting the application is the same as writing `-flag=true`
  - `-flag=true` and `-flag` are the same, you must use `-flag=false` for a `false` value
- Pre-existing variables
  - `var cfg config` where `config` is a struct
  - `flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")`
  - ... and so forth

## 3.2. Structured logging



## 3.3. Dependency injection

## 3.4. Centralized error handling

## 3.5. Isolating the application routes

