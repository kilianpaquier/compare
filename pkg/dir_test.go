package compare_test

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"

	compare "github.com/kilianpaquier/compare/pkg"
)

func TestDirs(t *testing.T) {
	testdata := filepath.Join("..", "testdata")

	t.Run("error_missing_dir", func(t *testing.T) {
		for _, tc := range []string{"actual", "expected"} {
			t.Run(tc, func(t *testing.T) {
				// Arrange
				testdata := filepath.Join(testdata, t.Name())
				expected := filepath.Join(testdata, "expected")
				actual := filepath.Join(testdata, "actual")

				// Act
				err := compare.Dirs(expected, actual)

				// Assert
				if err == nil {
					t.FailNow()
				}
				if !strings.Contains(err.Error(), "read dir") {
					t.Fatal(err)
				}
			})
		}
	})

	t.Run("error_missing_file", func(t *testing.T) {
		for _, tc := range []string{"actual", "expected"} {
			t.Run(tc, func(t *testing.T) {
				// Arrange
				testdata := filepath.Join(testdata, t.Name())
				expected := filepath.Join(testdata, "expected")
				actual := filepath.Join(testdata, "actual")

				// Act
				err := compare.Dirs(expected, actual)

				// Assert
				if err == nil {
					t.FailNow()
				}
				if !strings.Contains(err.Error(), fmt.Sprintf("missing file 'missing.txt' from %s", tc)) {
					t.Fatal(err)
				}
			})
		}
	})

	t.Run("error_not_equal", func(t *testing.T) {
		// Arrange
		testdata := filepath.Join(testdata, t.Name())
		expected := filepath.Join(testdata, "expected")
		actual := filepath.Join(testdata, "actual")
		files := []string{
			"file1.txt",
			"file2.txt",
			filepath.Join("subdir", "file1.txt"),
			filepath.Join("subdir", "file2.txt"),
		}

		// Act
		err := compare.Dirs(expected, actual)

		// Assert
		if err == nil {
			t.Fatal(err)
		}
		// small verification, there's no need for more since comparison result is handled by Golang diff library
		for _, file := range files {
			if !strings.Contains(err.Error(), fmt.Sprintf("diff %s %s", file, file)) {
				t.Error(fmt.Sprintf("missing diff for '%s'", file), err)
			}
		}
	})

	t.Run("success_equal", func(t *testing.T) {
		// Arrange
		testdata := filepath.Join(testdata, t.Name())
		expected := filepath.Join(testdata, "expected")
		actual := filepath.Join(testdata, "actual")

		// Act
		err := compare.Dirs(expected, actual)

		// Assert
		if err != nil {
			t.Fatal(err)
		}
	})
}
