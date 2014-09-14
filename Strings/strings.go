package strings

import (
	"strings"
)

func Index(s string, sep string, n int) int {
	i := strings.Index(s[n:], sep)
	if i > -1 {
		i += n
	}
	return i
}
