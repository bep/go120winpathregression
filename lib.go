package go120winpathregression

import "path/filepath"

func Clean(s string) string {
	return filepath.Clean(filepath.FromSlash(s))
}
