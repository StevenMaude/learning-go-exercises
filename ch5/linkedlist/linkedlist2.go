// Make use of the package container/list to create a (doubly) linked list.
// Push the values 1, 2 and 4 to the list and then print it.
// Create your own linked list implementation. And perform the same actions as above.
package main

import (
	"fmt"
)

type Item struct {
	previous *Item
	Value    int
	next     *Item
}

type List struct {
	first *Item
	last  *Item
}

func (l *List) PushBack(i int) {
	newItem := &Item{Value: i}
	if l.first == nil {
		// Only one item, which we've just added.
		l.first, l.last = newItem, newItem
	} else {
		// We need to make this the last item, and add links to
		// this newItem going back to the previous last, and the
		// previous last going forward to this newItem.
		newItem.previous = l.last
		newItem.previous.next = newItem
		l.last = newItem
	}
}

func (l *List) Front() *Item {
	return l.first
}

func (e *Item) Next() *Item {
	return e.next
}

func main() {
	l := &List{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(4)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
