# 4. Database-driven responses

## 4.1. Setting up MySQL

- Install and setup MysQL:

```bash
brew install mysql
brew services start mysql
mysql -u root -p
```

- Create the database and snippets table:

```sql
-- Create a new UTF-8 `snippetbox` database.
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
-- Switch to using the `snippetbox` database.
USE snippetbox;

-- Create a `snippets` table.
CREATE TABLE snippets (
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT, title VARCHAR(100) NOT NULL,
content TEXT NOT NULL,
created DATETIME NOT NULL,
expires DATETIME NOT NULL
);
-- Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);
```

- Add dummy records:

```sql
-- Add some dummy records (which we'll use in the next couple of chapters).
INSERT INTO snippets (title, content, created, expires) VALUES (
'An old silent pond',
'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō', UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'Over the wintry forest',
'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki', UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);
INSERT INTO snippets (title, content, created, expires) VALUES (
'First autumn morning',
'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo', UTC_TIMESTAMP(),
DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
```

- Create a new user for the web application:

```sql
CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE, DELETE ON snippetbox.* TO 'web'@'localhost'; -- Important: Make sure to swap 'pass' with a password of your own choosing. ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
```

- And test that you can access the database with this user:

```bash
mysql -D snippetbox -u web -p
```

```sql
SELECT id, title, expires FROM snippets;
```

## 4.2. Installing a database driver

```bash
go get github.com/go-sql-driver/mysql@v1
```


## 4.3. Modules and reproducible builds

- `go.mod` require lines tell the Go command exactly which version of a package should be used when you run a command like `go run`, `go test` or `go build`
- `// indirect` annotation indicates that a package doesn't directly appear in any `import` statement in your codebase (allows it to stay without causing an error, or a tool automatically removing the line)
- `go.sum` contains the cryptographic checksums representing the content of the required packages
- `go mod verify` verifies the checksums of the downloaded package on your machine, matching the entries in `go.sum`
- Upgrading packages
  - To upgrade to latest available _minor or patch release_ of a package, you can simply run `go get` like so:
    - `$ go get github.com/foo/bar`
  - Alternatively, upgrade to a specific version, use the `@version` suffix:
    - `$ go get github.com/foo/bar/@v1.2.3`
  - The `-u` flag of `go get` will upgrade the package _and all its dependencies to their latest versions_
    - `$ go get -u github.com/foo/bar`
  - Listing upgradable packages:
    - `$ go list -u -f '{{if (and (not (or .Main .Indirect)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all`
  - Removing unused packages:
    - There are two methods here
    - Run `go get` and postfix the package path with `@none`:
      - `$ go get github.com/foo/bar@none`
    - Or use `go mod tidy`:
      - `$ go mod tidy`


## 4.4. Creating a database connection pool

```go
// The sql.Open() function initializes a new sql.DB value, which is essentially
// a pool of database connections.
db, err := sql.Open("mysql", "web:pass@/snippetbox?parseTime=true")
if err != nil {
	...
}
````

- `sql.Open()`
  - first parameter: _driver name_
  - second parameter is the _data source name_ (DSN) which describes how to connect to your database
  - `parseTime=true` is a _driver-specific_ parameter which converts SQL `TIME` and `DATE` fields to Go `time.Time` values
  - returns a `sql.DB` value
    - a _pool of many connections_
    - Go manages the connections in this pool as needed, automatically opening and closing connections to the databnase via the driver
    - connection pool is safe for concurrent access
    - intended to be long-lived
    - actual connections to the database are established lazily, as and when needed for the first time
    - `db.Ping()` is used to create a connection and check that it's working
- import `_ "github.com/go-sql-driver/mysql"`
  - The underscore (`_`) is used when the code doesn't actually use anything in the package
  - The reason for importing is so that the `init()` function runs
  - Standard practice for most of Go's SQL drivers


## 4.5. Designing a database model

- The `internal` directory is being used to hold ancillary non-application-specific code which could potentially be re-used
- Benefits of this structure
  - Clean separation of concerns
    - Database logic not tied to handlers
    - Easier to write tight, focused, unit tests in the future
  - A custom `SnippetModel` allows us to
    - Make our model a single, neatly-encapsulated component
    - Can easily be initialized and passed to our handlers as a dependency
    - Allows for more maintable, testable code
  - Because model actions are defined as methods on the `SnippetModel` type
    - Opportunity to create an _interface_ and mock it for unit testing purposes
  - Total control over which database is used at runtime


## 4.6. Executing SQL statements

- Go provides three different methods for executing database queries:
  - `DB.Query()` is used for `SELECT` queries which return multiple rows
  - `DB.QueryRow()` is used for `SELECT` queries which return a single rows
  - `DB.Exec()` is used for statement which don't return rows (like `INSERT` and `DELETE`)
- `sql.Result` type returned by `DB.Exec()`:
  - `LastInsertId()` returns the integer (an `int64`) generated by the database in response to a command
    - Typically from an "auto-increment" column when inserting a row
    - Not all drivers support this
  - `RowsAffected()` returns the number of rows (as `int64`) affected by the statement
    - Not all drivers support this -- specifically, PostgreSQL does not support
- Placeholder parameters
  - `?` acts as placeholders
  - Help to avoid SQL injection attacks
  - `DB.Exec()` works in three steps
    1. Creates a new prepared statement on the database using the provided SQL statement. Database parses and compiles the statement, then stores it ready for execution
    2. In a second separate step, passes the parameter values to the database. DB then executes the prepared statement using these parameters. DB treats parameters as pure data, cannot change the _intent_ of the statement
    3. Closes (or _deallocates_) the prepared statement on the database
  - Note that parameter syntax differs depending on database
    - MySQL, SQL Server, SQLite all use `?` as the notation
    - PostgreSQL uses `$1`, `$2`, etc.


## 4.7. Single-record SQL queries

- `row.Scan()` will automatically convert the raw output from the SQL database to the required native Go types:
  - `CHAR`, `VARCHAR`, and `TEXT` map to `string`
  - `BOOLEAN` maps to `bool`
  - `INT` maps to `int`; `BIGINT` maps to `int64`
  - `DECIMAL` and `NUMERIC` map to `float`
  - `TIME`, `DATE`, and `TIMESTAMP` map to `time.Time`
  - MySQL driver quirk: need to use `parseTime=true` parameter in our DSN to force it to convert `TIME` and `DATE` fields to `time.Time`, otherwise they come over as `[]byte` slices
- Checking for specific errors:
  - `errors.Is()` checks whether an error matches a specific value
    - Works by _unwrapping_ errors as necessary before checking for a match
    - `errors.As()` checks for a specific type
- Shorthand single-record queries

```go
func (m *SnippetModel) Get(id int) (Snippet, error) {
	var s Snippet

	err := m.DB.QueryRow("SELECT ...", id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}
	
	return s, nil
}
```


## 4.8. Multiple-record SQL queries

- Closing a resultset with `defer rows.Close()` is critical
  - As long as resultset is open it will keep the underlying database connection open
  - If something goes wrong in the method where rows were returned, it can rapidly lead to all the connections in your pool being used up



## 4.9. Transactions and other details

- The database/sql package
  - Provides a standard interface between your Go application and the world of SQL databases
  - There _are_ some idiosyncrasies in the way that different drivers and databases operate
- Verbosity
  - Verbose, but non-magical
  - Can understand and control exactly what's going on
  - Consider `jmoiron/sqlx` package if verbosity really starts annoying you
- Managing null values
  - Go doesn't manage `NULL` values well
  - Set `NOT NULL` constraints on your database columns ... avoid `NULL` altogether
- Working with transactions
  - Guarantee the same connection is used by wrapping multiple statements in a transaction:
  
```go
type ExampleModel struct {
  DB *sql.DB
}

func (m *ExampleModel) ExampleTransaction() error {
	// Calling the Begin() method on the connection pool creates a new sql.Tx
	// value, which represents the in-progress database transaction.
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	
	// Defer a call to tx.Rollback() to ensure it is always called before the
	// function returns. If the transaction succeeds it will already be
	// committed by the time tx.Rollback() is called, making tx.Rollback() a
	// no-op. Otherwise, in the event of an error, tx.Rollback() will roll back
	// the changes before the function returns.
	defer tx.Rollback()
	
	// Call Exec() on the transaction, passing in your statement and any
	// parameters. It's important to notice that tx.Exec() is called on the
	// transaction value just created. NOT the connection pool. Although we're
	// using tx.Exec() here you can also use tx.Query() and tx.QueryRow() in
	// exactly the same way.
}
```

  - You must _always_ call either `Rollback()` or `Commit()` before your function returns
    - Otherwise connection will stay open and not be returned by the connection pool
    - This can lead to hitting your maximum connection limit/running out of resources
    - Simplest way is to use `defer tx.Rollback()`
  - Transactions are useful for executing multiple SQL statements as a _single atomic action_
    - Use `tx.Rollback()` and the transaction ensures:
      - _All_ statements are executed successfully; or
      - _No_ statements are eecuted and the database remains unchanged
- Prepared statements
  - `Exec()`, `Query()`, and `QueryRow()` automatically create prepared statements
  - May be inefficient ... better approach is to use `DB.Prepare()` to create a prepared statement once, and reuse that instead
    - Particularly true for complex SQL statements with multiple `JOINS`
  - Basic pattern:
  
```go
// We need somewhere to store the prepared statement for the lifetime of our
// web application. A neat way is to embed it in the model alongside the
// connection pool.
type ExampleModel struct {
	DB *sql.DB
	InsertStmt *sql.Stmt
}

// Create a constructor for the model, in which we set up the prepared
// statement.
func NewExampleModel(db *sql.DB) (*ExampleModel, error) {
	// Use the Prepare method to create a new prepared statement for the
	// current connection pool. This returns a sql.Stmt value which represents
	// the prepared statement.
	insertStmt, err := db.Prepare("INSERT INTO ...")
	if err != nil {
		return nil, err
	}
	
	// Store it in our ExampleModel struct, alongside the connection pool.
	return &ExampleModel{DB: db, InsertStmt: insertStmt}, nil
}

// Any methods implemented against the ExampleModel struct will have access to
// the prepared statement.
func (m *ExampleModel) Insert(args...) error {
	// We then need to call Exec directly against the prepared statement, rather
	// than against the connection pool. Prepared statement also support the
	// Query and QueryRow methods.
	_, err := m.InsertStmt.Exec(args...)
	
	return err
}

// In the web application's main function we will need to initialize a new
// ExampleModel struct using the constructor function.
func main() {
	db, err := sql.Open(...)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()

	// Use the constructor function to create a new ExampleModel struct.
	exampleModel, err := NewExampleModel(db)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	
	// Defer a call to Close() on the prepared statement to ensure that it is
	// properly closed before our main function terminates.
	defer exampleModel.InsertStmt.Close()
}
```

  - Prepared statements exist on _database connections_
    - Because Go uses a pool of _many database connections_, when you first use a prepared statement it will remember which connection in the pool was used
    - Next time prepared statement will attempt to use that same database connection
    - Heavier load, possibly that a large number of prepared statements will be created on multiple connections
      - Can lead to statements being prepared and re-prepared more often than you'd expect
      - Or even hitting server-side limits (MySQL: 16,382 prepared statements)
    - This code is more complicated than not using prepared statements
      - Measure accordingly
