package compare

// Error implements error and contains specific diffs between two files.
type Error struct{ diffs []byte }

var _ error = (*Error)(nil) // ensure interface is implemented

// Error implements error.
func (c *Error) Error() string {
	return string(c.diffs)
}
