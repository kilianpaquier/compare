package compare_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/kilianpaquier/compare/internal/testutils"
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
				testutils.Error(testutils.Require(t), err)
				testutils.Contains(t, err.Error(), "read dir")
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
				testutils.Error(testutils.Require(t), err)
				testutils.Contains(t, err.Error(), "missing file 'missing.txt' from "+tc)
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
		testutils.Error(testutils.Require(t), err)
		// small verification, there's no need for more since comparison result is handled by Golang diff library
		for _, file := range files {
			testutils.Contains(t, err.Error(), fmt.Sprintf("diff %s %s", file, file))
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
		testutils.NoError(t, err)
	})
}
