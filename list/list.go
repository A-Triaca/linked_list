package list

import (
	"bytes"
	"errors"
	"fmt"
)

type Node[V comparable] struct {
	value V
	next  *Node[V]
}

type List[V comparable] struct {
	head   *Node[V]
	length int
}

func (l *List[V]) Insert(value V) {
	n := Node[V]{
		value: value,
	}
	if l.head == nil {
		l.head = &n
		l.length += 1
		return
	}

	cur := l.head

	for cur.next != nil {
		cur = cur.next
	}

	cur.next = &n
	l.length += 1
}

func (l *List[V]) Remove(value V) error {
	cur := l.head
	var prev *Node[V] = nil
	if cur == nil {
		return errors.New("empty list")
	}
	for ok := true; ok; ok = cur != nil {
		if cur.value == value {
			if prev == nil {
				l.head = cur.next
			} else {
				prev.next = cur.next
			}
			l.length -= 1
			return nil
		}
		prev = cur
		cur = cur.next
	}
	return errors.New("value not in list")
}

func (l *List[V]) String() string {
	var out bytes.Buffer

	out.WriteString("(")

	cur := l.head
	if cur != nil {
		for {
			out.WriteString(fmt.Sprintf("%v", cur.value))
			if cur.next == nil {
				break
			}
			out.WriteString(", ")
			cur = cur.next
		}
	}
	out.WriteString(")")

	return out.String()
}
