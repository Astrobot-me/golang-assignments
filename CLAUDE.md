# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repo is

A personal Go learning repository: phase-wise practice assignments, one directory per phase. Each exercise file states its task in a comment header (concepts covered + requirements) with `// TODO: implement` markers; the point is for the repo owner to fill in the implementations. Theory notes live externally at https://maxthirty3.notion.site/golang (linked from README.md).

There is **no root go.mod** — commands like `go build ./...` do not work from the repo root. Each phase directory works differently:

## Directory layout and how to run things

### `medium-problems/` and `hard-problems/` — standalone scripts
Every `.go` file is `package main` with its own `main()` that prints demo output for manual verification. Because each file has a `main()`, the files in a directory do not compile together — run one file at a time:

```
cd medium-problems
go run m1_ctof.go
```

There are no tests here; correctness is checked by eyeballing the printed output against the expectations described in the file's header comment.

### `functions/`, `struct,interfaces & pointers/`, and `concurrency/` — modules with tests
Each of these directories is its own Go module (`module practice`, go 1.22). All files are `package practice`, and every exercise `NN_name.go` has a matching table-driven `NN_name_test.go`. Run from inside the directory:

```
cd functions
go test ./...                      # all tests in the phase
go test -run TestFilterSlice ./... # single exercise
go test -run 'TestFilterSlice/evens' ./...  # single table case
go vet ./...
```

Note: the `struct,interfaces & pointers` directory name contains a comma, spaces, and `&` — always quote it in shell commands.

The `concurrency/` tests read from exercise-returned channels via timeout-guarded helpers in `helpers_test.go`, so unimplemented or deadlocked solutions fail fast instead of hanging. `go test -race` is the real verifier for these exercises, but on this Windows machine it currently fails with "-race requires cgo" — a C toolchain (e.g. MSYS2/mingw-w64 gcc) plus `CGO_ENABLED=1` is needed first.

### `files.zip`
Each phase directory contains a `files.zip` holding the original (unsolved) exercise files. Don't delete or modify it.

## Conventions

- Do not rewrite the task-description comment headers or the test files — they define the exercises. Work happens in the solution stubs beneath the TODOs.
- Tests are table-driven, named `Test<Function>` with subtests per case; follow that pattern if adding exercises.
- Standard Go style: `gofmt` formatting, standard library only (no external dependencies anywhere in the repo).
- Commit subjects are short and imperative.
