# compare <!-- omit in toc -->

<p align="center">
  <img alt="GitHub Release" src="https://img.shields.io/github/v/release/kilianpaquier/compare?include_prereleases&sort=semver&style=for-the-badge">
  <img alt="GitHub Issues" src="https://img.shields.io/github/issues-raw/kilianpaquier/compare?style=for-the-badge">
  <img alt="GitHub License" src="https://img.shields.io/github/license/kilianpaquier/compare?style=for-the-badge">
  <img alt="Coverage" src="https://img.shields.io/codecov/c/github/kilianpaquier/compare/main?style=for-the-badge">
  <img alt="Go Version" src="https://img.shields.io/github/go-mod/go-version/kilianpaquier/compare/main?style=for-the-badge&label=Go+Version">
  <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/kilianpaquier/compare?style=for-the-badge">
</p>

---

- [How to use ?](#how-to-use-)
- [Documentation](#documentation)

## How to use ?

```sh
go get github.com/kilianpaquier/compare
```

## Documentation

Can be found here in a better format: https://pkg.go.dev/github.com/kilianpaquier/compare.

```go
/*
Package compare exposes two functions to compare directories and files between them.

Returned errors can be easily printed to get a easy visual on diffs.

All diffs are done with Golang internal diff library: https://github.com/golang/go/blob/master/src/internal/diff/diff.go

Kindly exported and maintained by go-internal: https://github.com/rogpeppe/go-internal/blob/master/diff/diff.go

Diffs are under the form of:

	diff path/to/expected.txt path/to/actual.txt
	--- path/to/expected.txt
	+++ path/to/actual.txt
	@@ -1,1 +1,1 @@
	-Some text value that should be equal to the expected one.
	\ No newline at end of file
	+Some text value that should be equal to the expected one
	\ No newline at end of file

Examples:

	func main() {
		expected := "path/to/expected.txt"
		actual := "path/to/actual.txt"

		err := compare.Files(expected, actual)
		// handle err
	}

	func main() {
		expected := "path/to/expected"
		actual := "path/to/actual"

		err := compare.Dirs(expected, actual)
		// handle err
	}
*/
package compare
```
