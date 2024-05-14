package list

import (
	"bytes"
	"fmt"
)

type Node struct {
	value int64
	next  *Node
}

type List struct {
	head   *Node
	length int
}

func (l *List) Insert(value int64) {
	n := Node{
		value: value,
	}
	if l.head == nil {
		l.head = &n
		return
	}

	cur := l.head

	for cur.next != nil {
		cur = cur.next
	}

	cur.next = &n
	l.length += 1
}

func (l *List) String() string {
	var out bytes.Buffer

	out.WriteString("(")

	cur := l.head
	for {
		out.WriteString(fmt.Sprintf("%d", cur.value))
		if cur.next == nil {
			break
		}
		out.WriteString(", ")
		cur = cur.next
	}

	out.WriteString(")\n")

	return out.String()
}
