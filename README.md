# golang_course

Learning Go — practice programs.

## Project layout

```
.
├── profit_calculator.go   # standalone example at the repo root
└── course1/
    ├── bank.go            # banking program (package main)
    ├── math_package_guide.go  # standalone guide to the math package (package main)
    └── go.mod
```

## Running the programs

### `course1/bank.go`

```bash
cd course1
go run bank.go
```

> **Don't use `go run bank`.** `go run` expects an *import path* or `.go` files,
> not a directory/binary name. Passing `bank` makes Go look for a package called
> `bank` in the standard library and fail with:
>
> ```
> package bank is not in std (/usr/local/go/src/bank)
> ```
>
> Note also that `course1/bank` is a leftover compiled binary (from `go build`).
> Run it directly with `./bank` if you want the compiled version.

### Heads up: one `main` per directory

`bank.go` and `math_package_guide.go` both live in `course1/` and both declare
`package main` with `func main()`. A Go directory is a single package and may
only have one `main`, so `go run .` fails here with `main redeclared in this
block`. Run a single file (`go run bank.go`) until the guide is moved into its
own folder.
