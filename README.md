# gator 🐊

`gator` is a command-line RSS feed aggregator written in Go.
It allows users to follow feeds, fetch posts, and browse recent entries directly from the terminal.

---

## Requirements

Before running `gator`, make sure you have the following installed:

* **Go** (1.21+ recommended)
* **PostgreSQL**

### Install Go

Follow the official instructions for your OS:
[https://go.dev/doc/install](https://go.dev/doc/install)

Verify installation:

```
go version
```

### Install PostgreSQL

Install PostgreSQL using your system’s package manager or from:
[https://www.postgresql.org/download/](https://www.postgresql.org/download/)

Verify installation:

```
psql --version
```

Make sure PostgreSQL is running and you have a database created for `gator`.

---

## Installation

`gator` is a statically compiled Go binary. Once installed, it does **not** require the Go toolchain to run.

Install `gator` using:

```
go install github.com/ElitistNoob/gator@latest
```

Ensure your Go bin directory is in your `PATH`:

```
export PATH="$PATH:$(go env GOPATH)/bin"
```

After this, you should be able to run:

```
gator
```

---

## Configuration

`gator` uses a JSON config file located at:

```
~/.gatorconfig.json
```

### Example config file

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": "yourname"
}
```

### Fields

* `db_url` – PostgreSQL connection string
* `current_user_name` – the active user for CLI commands

---

## Usage

### Create a user

```
gator register <username>
```

### Log in as a user

```
gator login <username>
```

### Add a feed

```
gator addfeed <feed-url>
```

### Follow a feed

```
gator follow <feed-url>
```

### Fetch posts

```
gator agg
```

### Browse recent posts

```
gator browse [limit]
```

Example:

```
gator browse 5
```

---

## Development vs Production

During development, you may use:

```
go run .
```

This is **only for development**.

For normal usage, run the compiled binary:

```
gator
```

After running `go build` or `go install`, the program is fully self-contained and does **not** require Go to be installed.

---

## Repository

Source code is hosted on GitHub:

[https://github.com/](https://github.com/ElitistNoob/gator)

---

## License

MIT
