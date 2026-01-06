# compare <!-- omit in toc -->

<div align="center">
  <img alt="GitLab Release" src="https://img.shields.io/gitlab/v/release/kilianpaquier%2Fcompare?gitlab_url=https%3A%2F%2Fgitlab.com&include_prereleases&sort=semver&style=for-the-badge">
  <img alt="GitLab Issues" src="https://img.shields.io/gitlab/issues/open/kilianpaquier%2Fcompare?gitlab_url=https%3A%2F%2Fgitlab.com&style=for-the-badge">
  <img alt="GitLab License" src="https://img.shields.io/gitlab/license/kilianpaquier%2Fcompare?gitlab_url=https%3A%2F%2Fgitlab.com&style=for-the-badge">
  <img alt="GitLab CICD" src="https://img.shields.io/gitlab/pipeline-status/kilianpaquier%2Fcompare?gitlab_url=https%3A%2F%2Fgitlab.com&branch=main&style=for-the-badge">
  <img alt="Go Version" src="https://img.shields.io/gitlab/go-mod/go-version/kilianpaquier/compare?style=for-the-badge">
  <img alt="Go Report Card" src="https://goreportcard.com/badge/gitlab.com/kilianpaquier/compare?style=for-the-badge">
</div>

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

	func main() {
		expected := "expected content (only useful on multiline contents with the desire to get a pretty diff)"
		actual := "actual content (only useful on multiline contents with the desire to get a pretty diff)"

		err := compare.Contents(expected, actual)
		// handle err
	}
*/
package compare
```
