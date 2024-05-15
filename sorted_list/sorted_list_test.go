package sorted_list

import (
	"testing"
)

func TestInserting(t *testing.T) {
	tests := []struct {
		input  []int64
		output string
		length int
	}{
		{[]int64{1}, "(1)", 1},
		{[]int64{1, 3, 2, 4}, "(1, 2, 3, 4)", 4},
		{[]int64{4, 3, 2, 1}, "(1, 2, 3, 4)", 4},
		{[]int64{4, 3, 2, 1, 5}, "(1, 2, 3, 4, 5)", 5},
	}

	for _, test := range tests {
		l := List{}
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
		input  []int64
		remove int64
		output string
		err    string
		length int
	}{
		{[]int64{1}, 1, "()", "", 0},
		{[]int64{1, 2, 3}, 2, "(1, 3)", "", 2},
		{[]int64{1, 2}, 2, "(1)", "", 1},
		{[]int64{1, 2}, 1, "(2)", "", 1},
		{[]int64{1}, 2, "(1)", "value not in list", 1},
		{[]int64{}, 2, "()", "empty list", 0},
	}

	for _, test := range tests {
		l := List{}
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
