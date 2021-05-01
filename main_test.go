package kmp

import (
	"reflect"
	"testing"
)

func TestMatching(t *testing.T) {
	tables := []struct {
		text    string
		pattern string
		exp     []int
	}{
		{"hello ello el", "ell", []int{1, 6}},
		{"aaaabcabaabba", "ab", []int{3, 6, 9}},
		{"a", "a", []int{0}},
		{"a", "b", []int{}},
	}

	for _, table := range tables {
		res := FindMatches(table.text, table.pattern)

		if !reflect.DeepEqual(res, table.exp) {
			t.Fatalf("Expected %v, but got %v\n", table.exp, res)
		}
	}
}
