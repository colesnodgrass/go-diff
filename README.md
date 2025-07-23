# go-diff

Wrapper around `go-cmp` that adds color output.

Includes three helper functions:
- diff.Diff - returns a `go-cmp` string with appropriate colored output
- diff.Error - intended for use in tests, passes colored output to `testing.T.Error`
- diff.Fatal - intended for use in tests, passes colored output to `testing.T.Fatal`
