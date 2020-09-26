package hw04_lru_cache //nolint:golint,stylecheck

type LinkedListItem struct {
	Value interface{}
	Prev  *LinkedListItem
	Next  *LinkedListItem
}

type LinkedList struct {
	Head   *LinkedListItem
	Tail   *LinkedListItem
	Length int
}

type List interface {
	Len() int
	Front() *LinkedListItem
	Back() *LinkedListItem
	PushFront(v interface{}) *LinkedListItem
	PushBack(v interface{}) *LinkedListItem
	Remove(i *LinkedListItem)
	MoveToFront(i *LinkedListItem)
}

func (l *LinkedList) Len() int {
	return l.Length
}

func (l *LinkedList) Front() *LinkedListItem {
	return l.Head
}

func (l *LinkedList) Back() *LinkedListItem {
	return l.Tail
}

func (l *LinkedList) PushFront(v interface{}) *LinkedListItem {
	item := LinkedListItem{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}

	l.Length++

	if l.Head != nil {
		l.Head.Prev = &item
		item.Next = l.Head
		l.Head = &item
	} else {
		l.Head = &item
		l.Tail = &item
	}

	return &item
}

func (l *LinkedList) PushBack(v interface{}) *LinkedListItem {
	item := LinkedListItem{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}

	l.Length++

	if l.Tail != nil {
		l.Tail.Next = &item
		item.Prev = l.Tail
		l.Tail = &item
	} else {
		l.Head = &item
		l.Tail = &item
	}

	return &item
}

func (l *LinkedList) Remove(i *LinkedListItem) {
	l.Length--

	if l.Length == 0 {
		l.Head = nil
		l.Tail = nil
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next

		if l.Tail == i {
			l.Tail = i.Prev
		}
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev

		if l.Head == i {
			l.Head = i.Next
		}
	}
}

func (l *LinkedList) MoveToFront(i *LinkedListItem) {
	if i.Prev == nil {
		return
	}

	i.Prev.Next = i.Next

	if l.Tail != i {
		i.Next.Prev = i.Prev
	} else {
		l.Tail = i.Prev
	}

	i.Next = l.Head
	i.Prev = nil
	l.Head = i
}

func NewList() *LinkedList {
	var _ List = (*LinkedList)(nil)
	return &LinkedList{}
}
