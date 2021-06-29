package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Prev  *ListItem
	Next  *ListItem
}

type Lists struct {
	head *ListItem
	tail *ListItem
	len  int
}

func (l *Lists) Len() int         { return l.len }
func (l *Lists) Front() *ListItem { return l.head }
func (l *Lists) Back() *ListItem  { return l.tail }

func (l *Lists) PushFront(v interface{}) *ListItem {
	newElem := ListItem{v, nil, l.head}

	if l.len == 0 {
		l.tail = &newElem
	} else {
		l.head.Prev = &newElem
	}
	l.head = &newElem
	l.len++
	return l.head
}

func (l *Lists) PushBack(v interface{}) *ListItem {
	newElem := ListItem{v, l.tail, nil}

	if l.len == 0 {
		l.head = &newElem
	} else {
		l.tail.Next = &newElem
	}

	l.tail = &newElem
	l.len++

	return l.tail
}

func (l *Lists) Remove(i *ListItem) {
	if i.Prev == nil && i.Next == nil {
		l.head = nil
		l.tail = nil
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.head = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}

	i.Value = nil
	i.Prev = nil
	i.Next = nil
	l.len--
}

func (l *Lists) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	frontElem := ListItem{i.Value, nil, l.head}

	l.head.Prev = &frontElem
	i.Prev.Next = i.Next

	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}

	l.head = &frontElem
}

func NewList() *Lists {
	l := Lists{nil, nil, 0}
	return &l
}
