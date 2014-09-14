package strings

import (
	"testing"
)

var s = "go gopher, go"

func Test_IndexShouldReturnMinusOneIfSepNotFound(t *testing.T) {
	i := Index(s, "abc", 0)

	if i != -1 {
		t.Errorf("Index should be -1 but is %d", i)
	}
}

func Test_IndexShouldReturnCorrectIndexWithStartIndex(t *testing.T) {
	i := Index(s, "go", 1)

	if i != 3 {
		t.Errorf("Index should be 3 but is %d", i)
	}
}
