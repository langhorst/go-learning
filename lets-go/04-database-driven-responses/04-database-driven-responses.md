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

## 4.5. Designing a database model

## 4.6. Executing SQL statements

## 4.7. Single-record SQL queries

## 4.8. Multiple-record SQL queries

## 4.9. Transactions and other details
