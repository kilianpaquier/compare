package compare

import (
	"bytes"
	"fmt"
	"os"

	"github.com/rogpeppe/go-internal/diff"
)

// Contents compares expected and actual bytes.
//
// It returns an error if there're differences.
// Differences are wrapped in Error implementation of error.
//
// Carriage character '\r' is removed from contents before comparison
// to get success results when comparing the same file between windows and linux.
func Contents(expected, actual []byte) error {
	return contents("expected", "actual", expected, actual)
}

// Files compares expected filepath and actual filepath contents.
//
// It returns an error if there're differences.
// Differences are wrapped in Error implementation of error.
//
// Carriage character '\r' is removed from contents before comparison
// to get success results when comparing the same file between windows and linux.
func Files(expected, actual string) error {
	expectedContent, err := os.ReadFile(expected)
	if err != nil {
		return fmt.Errorf("read expected file: %w", err)
	}

	actualContent, err := os.ReadFile(actual)
	if err != nil {
		return fmt.Errorf("read actual file: %w", err)
	}

	return contents(expected, actual, expectedContent, actualContent)
}

// contents compares bytes directly.
//
// It returns an error if there're differences.
// Differences are wrapped in Error implementation of error.
//
// Carriage character '\r' is removed from contents before comparison
// to get success results when comparing the same file between windows and linux.
func contents(ename, aname string, expected, actual []byte) error { //nolint:revive
	diffs := diff.Diff(ename, bytes.ReplaceAll(expected, Carriage, []byte{}), aname, bytes.ReplaceAll(actual, Carriage, []byte{}))
	if len(diffs) > 0 {
		return &Error{diffs}
	}
	return nil
}
