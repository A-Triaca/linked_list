package list

import (
	"testing"
)

func TestInsertingInts(t *testing.T) {
	tests := []struct {
		input  []int
		output string
		length int
	}{
		{[]int{1}, "(1)", 1},
		{[]int{1, 2, 3, 4}, "(1, 2, 3, 4)", 4},
	}

	for _, test := range tests {
		l := List[int]{}
		for _, i := range test.input {
			l.Insert(i)
		}
		if l.String() != test.output {
			t.Fatalf("expected=%v. got=%v", test.output, l.String())
		}
		if l.length != test.length {
			t.Fatalf("expected length=%d. got length=%d", test.length, l.length)
		}
	}
}
func TestInsertingStrings(t *testing.T) {
	tests := []struct {
		input  []string
		output string
		length int
	}{
		{[]string{"a"}, "(a)", 1},
		{[]string{"a", "b", "c", "d"}, "(a, b, c, d)", 4},
	}

	for _, test := range tests {
		l := List[string]{}
		for _, i := range test.input {
			l.Insert(i)
		}
		if l.String() != test.output {
			t.Fatalf("expected=%v. got=%v", test.output, l.String())
		}
		if l.length != test.length {
			t.Fatalf("expected length=%d. got length=%d", test.length, l.length)
		}
	}
}

func TestRemoving(t *testing.T) {
	tests := []struct {
		input  []int
		remove int
		output string
		err    string
		length int
	}{
		{[]int{1}, 1, "()", "", 0},
		{[]int{1, 2, 3}, 2, "(1, 3)", "", 2},
		{[]int{1, 2}, 2, "(1)", "", 1},
		{[]int{1, 2}, 1, "(2)", "", 1},
		{[]int{1}, 2, "(1)", "value not in list", 1},
		{[]int{}, 2, "()", "empty list", 0},
	}

	for _, test := range tests {
		l := List[int]{}
		for _, i := range test.input {
			l.Insert(i)
		}

		ok := l.Remove(test.remove)

		if ok != nil && ok.Error() != test.err {
			t.Fatalf("expected error=%v. got error=%v", test.err, ok)
		}

		if l.String() != test.output {
			t.Fatalf("expected=%v. got=%v", test.output, l.String())
		}
		if l.length != test.length {
			t.Fatalf("expected length=%d. got length=%d", test.length, l.length)
		}
	}
}
