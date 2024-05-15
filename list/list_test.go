package list

import (
	"testing"
)

func TestInserting(t *testing.T) {
	tests := []struct {
		input  []int64
		output string
	}{
		{[]int64{1}, "(1)"},
		{[]int64{1, 2, 3, 4}, "(1, 2, 3, 4)"},
	}

	for _, test := range tests {
		l := List{}
		for _, i := range test.input {
			l.Insert(i)
		}
		if l.String() != test.output {
			t.Fatalf("expected=%v. got=%v", test.output, l.String())
		}
	}
}

func TestRemoving(t *testing.T) {
	tests := []struct {
		input  []int64
		remove int64
		output string
		err    string
	}{
		{[]int64{1}, 1, "()", ""},
		{[]int64{1, 2, 3}, 2, "(1, 3)", ""},
		{[]int64{1, 2}, 2, "(1)", ""},
		{[]int64{1, 2}, 1, "(2)", ""},
		{[]int64{1}, 2, "(1)", "value not in list"},
		{[]int64{}, 2, "()", "empty list"},
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
	}
}
