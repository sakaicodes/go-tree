# go-tree

go-tree is a small Go CLI that prints a directory structure in a tree format, similar to the Linux `tree` command.

## Features

- Recursive directory tree output
- Unicode tree connectors for readable structure
- Optional `-all` flag to include hidden and normally skipped directories
- Simple default skips for common noisy folders

## Requirements

- Go 1.26.1 or newer (based on `go.mod`)

## Install and Run

### Run directly

From the project root:

```bash
go run . <directory>
```

Example:

```bash
go run . .
```

### Build a binary

From the project root:

```bash
go build -o go-tree .
```

Then run:

```bash
./go-tree <directory>
```

## Usage

```text
go-tree [flags] <directory>
```

If you run without an argument, the program exits with:

```text
Usage: go-tree [directory]
```

### Flags

- `-all`: include everything (including hidden paths and paths that are skipped by default)

Examples:

```bash
go run . .
go run . -all .
./go-tree ./test
```

## Example Output

Command:

```bash
go run . .
```

Output:

```text
├── .gitignore
├── README.md
├── cmd
│   └── display_tree.go
├── go.mod
├── internal
│   ├── tree.go
│   └── utils.go
├── main.go
├── test
│   ├── subtest
│   │   └── subtest1.go
│   ├── test1.go
│   └── test2.go
└── tree
```

## Default Skip List

When `-all` is not provided, the program skips:

- `.git`
- `node_modules`
- `.DS_Store`

This list is defined in `internal/utils.go`.

## Project Structure

- `main.go`: CLI entrypoint and argument guard
- `cmd/display_tree.go`: argument and flag parsing, handoff to internal logic
- `internal/tree.go`: recursive tree printing logic
- `internal/utils.go`: default skip configuration

## Current Notes

- The tool expects a directory argument.
- Passing `.` prints from the current working directory.
- Errors from unreadable paths are surfaced by panic in the current implementation.

## License

No license file is currently included in this repository.
