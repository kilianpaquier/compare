package compare

import (
	"bytes"
	"errors"
	"fmt"
	"maps"
	"os"
	"path/filepath"
)

// Carriage represents the \r character in byte format.
var Carriage = []byte{13}

// Dirs compares expected an actual directories (and their subdirectories).
//
// It returns an join'ed slice of errors if there're differences.
// Differences are wrapped in Error implementation of error.
//
// Carriage character '\r' is removed from contents before comparison
// to get success results when comparing the same file between windows and linux.
func Dirs(expected, actual string) error {
	// read a directory by getting it's absolute path
	// and then iterating recursively to parse all relative files to dir.
	read := func(path string) (map[string][]byte, error) {
		abs, err := filepath.Abs(path)
		if err != nil {
			return nil, fmt.Errorf("absolute path: %w", err)
		}
		files, err := readDir(abs, abs)
		if err != nil {
			return nil, err
		}
		return files, nil
	}

	expectedFiles, err := read(expected)
	if err != nil {
		return err
	}

	actualFiles, err := read(actual)
	if err != nil {
		return err
	}

	// check all expected contents against actual contents
	var errs []error
	for file, expectedBytes := range expectedFiles {
		actualBytes, ok := actualFiles[file]
		if !ok {
			errs = append(errs, fmt.Errorf("missing file '%s' from actual", file))
			continue
		}
		errs = append(errs, contents(file, file, expectedBytes, actualBytes))
	}

	// check that there're no actual files that aren't present in expected files
	for file := range actualFiles {
		if _, ok := expectedFiles[file]; !ok {
			errs = append(errs, fmt.Errorf("missing file '%s' from expected", file))
		}
	}

	// Join takes care of removing nil errors and returning nil in case all errors were nil
	return errors.Join(errs...)
}

// readDir reads a given input directory (and its subdirectories)
// and returns a map with files path as keys and content (slice of bytes) as values.
func readDir(initialdir string, currentdir string) (map[string][]byte, error) {
	files := map[string][]byte{}

	entries, err := os.ReadDir(currentdir)
	if err != nil {
		return nil, fmt.Errorf("read dir: %w", err)
	}

	errs := make([]error, 0, len(entries))
	for _, entry := range entries {
		src := filepath.Join(currentdir, entry.Name())

		// handle directories
		if entry.IsDir() {
			sub, err := readDir(initialdir, src)
			if err != nil {
				errs = append(errs, err) // only case of error is if reading an entry fails
				continue
			}

			maps.Copy(files, sub)
			continue
		}

		// handle files
		content, err := os.ReadFile(src)
		if err != nil {
			errs = append(errs, fmt.Errorf("read file: %w", err))
			continue
		}

		abs, err := filepath.Abs(src)
		if err != nil {
			errs = append(errs, fmt.Errorf("absolute path: %w", err))
			continue
		}

		rel, err := filepath.Rel(initialdir, abs)
		if err != nil {
			errs = append(errs, fmt.Errorf("relative path: %w", err))
			continue
		}
		files[rel] = bytes.ReplaceAll(content, Carriage, []byte{})
	}
	return files, errors.Join(errs...)
}
