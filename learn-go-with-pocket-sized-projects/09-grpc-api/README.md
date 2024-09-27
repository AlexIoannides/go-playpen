# A Habits Tracker using gRPC

- defining Protobuf schema
  - auto-generate serialisation/de-serialisation Go code to enable native use of data objects in codebase
- defining gRCP servers
- embedding structs as a form of composition
- using `go:generate` metadata tagging to auto-generate code
- mocking test using `github.com/gojuno/minimock`
- using `errgroup.Group` as an alternative to `sync.WaitGroup`
- using `if testing.Short() { t.Skip() }` with `go test -short` to exclude long-running tests from executing
- using `context` from the std library to carry around deadlines, cancellation signals, and values across API boundaries
