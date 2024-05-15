package sorted_list

import (
	"bytes"
	"errors"
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
		l.length += 1
		return
	}

	cur := l.head
	var prev *Node
	for cur.next != nil && cur.value < value {
		prev = cur
		cur = cur.next
	}
	if cur.next == nil && cur.value < value {
		cur.next = &n
	} else if prev == nil {
		n.next = cur
		l.head = &n
	} else {
		prev.next = &n
		n.next = cur
	}
	l.length += 1
}

func (l *List) Remove(value int64) error {
	cur := l.head
	var prev *Node = nil
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

func (l *List) String() string {
	var out bytes.Buffer

	out.WriteString("(")

	cur := l.head
	if cur != nil {
		for {
			out.WriteString(fmt.Sprintf("%d", cur.value))
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
