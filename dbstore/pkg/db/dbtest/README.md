Cannot import test code from another package so helper functions have to go into proper code files rather than test files.

If it isn't imported from non-test code then it won't be built into the final binary.

example package ...test

[httptest example](https://golang.org/pkg/net/http/httptest/)
[zaptest example](https://godoc.org/go.uber.org/zap/zaptest)
