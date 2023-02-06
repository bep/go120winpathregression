package go120winpathregression

import (
	"path/filepath"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestTo(t *testing.T) {
	c := qt.New(t)
	c.Assert(Clean("////post/////"), qt.Equals, filepath.FromSlash("/post"))
}
