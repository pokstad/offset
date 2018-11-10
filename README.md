Offset is a simple tool for calculating the byte offset for a rune in a file at a specificed line and column number.

Offset is both usable as an executable:

```
# Install with go get

$ go get github.com/pokstad/offset

# Run with a file arg
$ offset -source test.go -line 3 -column 20
38

# Or omit file arg to provide STDIN:
$ offset -line 1 -column 2
hello
2
```

Or as a Go package:

[![GoDoc](https://godoc.org/github.com/pokstad/offset/offset?status.svg)](https://godoc.org/github.com/pokstad/offset/offset)
