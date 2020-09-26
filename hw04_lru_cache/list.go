package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Value interface{}
	Prev  *listItem
	Next  *listItem
}

type list struct {
	head   *listItem
	tail   *listItem
	length int
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *listItem {
	return l.head
}

func (l *list) Back() *listItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *listItem {
	item := listItem{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}

	l.length++

	if l.head != nil {
		l.head.Prev = &item
		item.Next = l.head
		l.head = &item
	} else {
		l.head = &item
		l.tail = &item
	}

	return &item
}

func (l *list) PushBack(v interface{}) *listItem {
	item := listItem{
		Value: v,
		Prev:  nil,
		Next:  nil,
	}

	l.length++

	if l.tail != nil {
		l.tail.Next = &item
		item.Prev = l.tail
		l.tail = &item
	} else {
		l.head = &item
		l.tail = &item
	}

	return &item
}

func (l *list) Remove(i *listItem) {
	l.length--

	if l.length == 0 {
		l.head = nil
		l.tail = nil
		return
	}

	if l.head != i {
		i.Prev.Next = i.Next

		if l.tail == i {
			l.tail = i.Prev
		}
	}

	if l.tail != i {
		i.Next.Prev = i.Prev

		if l.head == i {
			l.head = i.Next
		}
	}
}

func (l *list) MoveToFront(i *listItem) {
	if l.head == i {
		return
	}

	i.Prev.Next = i.Next

	if l.tail != i {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}

	i.Next = l.head
	i.Prev = nil
	l.head = i
}

func NewList() List {
	return &list{}
}
