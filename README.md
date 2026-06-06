# golang_course

Learning Go — practice programs.

## Project layout

```
.
├── profit_calculator.go   # standalone example at the repo root
└── course1/
    ├── bank.go            # banking program — main, balance file I/O
    ├── menu.go            # presentMenu() helper, same package main
    └── go.mod
```

## Running the programs

### `course1/` (the bank program)

```bash
cd course1
go run .
```

Use `go run .` — it compiles **every** `.go` file in the directory as one
package. The bank program is split across `bank.go` and `menu.go`, so they have
to be built together.

## Common gotchas

**`go run bank` →** `package bank is not in std (/usr/local/go/src/bank)`
`go run` expects an *import path* or `.go` files, not a directory/binary name.
Passing `bank` makes Go look for a package called `bank` in the standard library.
(If a compiled `bank` binary exists from `go build`, run it directly with
`./bank`.)

**`go run bank.go` →** `undefined: presentMenu`
Naming a single file compiles *only that file*. Functions defined in sibling
files of the same package (e.g. `presentMenu` in `menu.go`) won't be found. Use
`go run .` instead.

**`go run .` →** `main redeclared in this block`
A Go directory is a single package and may have only one `func main()`. If you
get this, two files in the directory both declare `main` — move one into its own
folder.
