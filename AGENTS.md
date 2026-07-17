# Repository Guidelines

## Project Structure & Module Organization

This is a Go repository. Keep source code organized by package, with one focused package per directory. Put executable entry points in `cmd/<app>/` as needed, reusable private code in `internal/`, and tests next to their implementation as `*_test.go`. Keep assets and test fixtures in the package that uses them.

Do not commit build artifacts, coverage output, local configuration, or editor-specific files.

## Build, Test, and Development Commands

- `go build ./...` — compiles every package.
- `go test ./...` — runs the complete test suite.
- `go test -race ./...` — checks concurrent code for races.
- `go vet ./...` — reports common correctness issues.
- `go fmt ./...` — formats Go source using the standard formatter.

Run commands from the repository root. To run an executable package locally, use `go run ./cmd/<app>` when that directory exists.

## Coding Style & Naming Conventions

Write idiomatic Go and always run `gofmt`; it determines indentation and layout. Use concise lowercase package names without underscores. Exported identifiers use `PascalCase`; unexported identifiers use `camelCase`. Use descriptive responsibility-based filenames, such as `handler.go`, `config.go`, and `user_test.go`.

Prefer small functions, explicit error handling, and standard-library APIs. Add useful operation context when wrapping errors, for example: `fmt.Errorf("load config: %w", err)`.

## Testing Guidelines

Use table-driven tests for related cases. Name tests `Test<Function>_<Scenario>` and benchmarks `Benchmark<Function>`. Tests must be deterministic and should not rely on network access, wall-clock time, or shared machine state. Add a regression test with each bug fix, then run `go test ./...`.

## Commit & Pull Request Guidelines

Use short, imperative commit subjects such as `Add config validation` or `Fix empty response handling`. Keep commits focused on one logical change. Pull requests should explain the behavior change, list validation performed, link relevant issues, and include screenshots or sample output for user-facing changes.
