package utils

import (
	"net/url"
	"strings"
)

func EncodePath(pathname string) string {
	paths := strings.Split(pathname, "/")
	for i, path := range paths {
		paths[i] = url.PathEscape(path)
	}
	return strings.Join(paths, "/")
}
