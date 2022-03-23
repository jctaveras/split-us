package internal

import (
	"path"
	"strings"
)

func ShiftPath(pattern string) (head, tail string) {
	pattern = path.Clean("/" + pattern)
	index := strings.Index(pattern[1:], "/") + 1

	if index <= 0 {
		return pattern[1:], "/"
	}

	return pattern[1:index], pattern[index:]
}
