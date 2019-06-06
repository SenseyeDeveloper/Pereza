package parser

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func getPackagePathByEnv(prefixes []string, pwd, filename string, isDir bool) (string, error) {
	if !filepath.IsAbs(filename) {
		filename = filepath.Join(pwd, filename)
	}

	for _, prefix := range prefixes {
		if rel := strings.TrimPrefix(filename, prefix); rel != filename {
			if isDir {
				return path.Clean(filepath.ToSlash(rel)), nil
			}

			return path.Dir(filepath.ToSlash(rel)), nil
		}
	}

	return "", fmt.Errorf("file '%v' is not in GOPATH", filename)
}

func prefixes(gopath string) []string {
	gopaths := strings.Split(gopath, string(filepath.ListSeparator))

	prefixes := make([]string, len(gopaths))

	for i, p := range gopaths {
		prefixes[i] = filepath.Join(p, "src") + string(filepath.Separator)
	}

	return prefixes
}
