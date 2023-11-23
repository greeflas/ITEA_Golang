package filesystem

import (
	"strings"
)

func BuildPath(parts ...string) string {
	return strings.Join(parts, "/")
}
