package compare_test

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"testing"

	"github.com/kilianpaquier/compare/internal/testutils"
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
				testutils.ErrorIs(t, err, fs.ErrNotExist)
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
		testutils.ErrorAs(testutils.Require(t), err, &ce)
		// small verification, there's no need for more since comparison result is handled by Golang diff library
		testutils.Contains(t, ce.Error(), fmt.Sprintf("diff %s %s", expected, actual))
	})

	t.Run("success", func(t *testing.T) {
		for _, tc := range []string{"empty", "equal"} {
			t.Run(tc, func(t *testing.T) {
				// Arrange
				testdata := filepath.Join(testdata, t.Name())
				expected := filepath.Join(testdata, "expected.txt")
				actual := filepath.Join(testdata, "actual.txt")

				// Act
				err := compare.Files(expected, actual)

				// Assert
				testutils.NoError(t, err)
			})
		}
	})
}
