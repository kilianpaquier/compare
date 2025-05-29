package compare_test

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	compare "github.com/kilianpaquier/compare/pkg"
)

func TestFiles(t *testing.T) {
	testdata := filepath.Join("..", "testdata")

	t.Run("missing", func(t *testing.T) {
		for _, tc := range []string{"expected", "actual"} {
			t.Run(tc, func(t *testing.T) {
				// Arrange
				testdata := filepath.Join(testdata, t.Name())
				expected := filepath.Join(testdata, "expected.txt")
				actual := filepath.Join(testdata, "actual.txt")

				// Act
				err := compare.Files(expected, actual)

				// Assert
				if !errors.Is(err, fs.ErrNotExist) {
					t.Fatal(err)
				}
			})
		}
	})

	t.Run("error_not_equal", func(t *testing.T) {
		// Arrange
		testdata := filepath.Join(testdata, t.Name())
		expected := filepath.Join(testdata, "expected.txt")
		actual := filepath.Join(testdata, "actual.txt")

		// Act
		err := compare.Files(expected, actual)

		// Assert
		ce := &compare.Error{}
		if !errors.As(err, &ce) {
			t.Fatal(err)
		}
		// small verification, there's no need for more since comparison result is handled by Golang diff library
		if !strings.Contains(ce.Error(), fmt.Sprintf("diff %s %s", expected, actual)) {
			t.Fatal(err)
		}
	})

	t.Run("success", func(t *testing.T) {
		for _, tc := range []string{"empty", "equal"} {
			t.Run(tc, func(t *testing.T) {
				// Arrange
				testdata := filepath.Join(testdata, t.Name())
				expected := filepath.Join(testdata, "expected.txt")
				actual := filepath.Join(testdata, "actual.txt")

				// Act & Assert
				err := compare.Files(expected, actual)
				if err != nil {
					t.Fatal(err)
				}
			})
		}
	})
}
