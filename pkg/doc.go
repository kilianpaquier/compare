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
