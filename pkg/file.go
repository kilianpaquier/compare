package compare

import (
	"bytes"
	"os"

	"github.com/rogpeppe/go-internal/diff"
)

// Files compares expected filepath and actual filepath contents.
//
// It returns an error if there're differences.
// Differences are wrapped in CompareError implementation of error.
//
// Carriage character '\r' is removed from contents before comparison
// to get success results when comparing the same file between windows and linux.
func Files(expected, actual string) error {
	expectedContent, _ := os.ReadFile(expected)
	actualContent, _ := os.ReadFile(actual)

	diffs := diff.Diff(expected, bytes.ReplaceAll(expectedContent, carriage, []byte{}), actual, bytes.ReplaceAll(actualContent, carriage, []byte{}))
	if len(diffs) > 0 {
		return &Error{diffs}
	}
	return nil
}
