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
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
	len   int
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	var newElement ListItem
	newElement.Value = v
	if l.len > 0 {
		newElement.Next = l.first
		l.first.Prev = &newElement
		l.first = &newElement
	} else {
		l.first, l.last = &newElement, &newElement
	}
	l.len++

	return l.first
}

func (l *list) PushBack(v interface{}) *ListItem {
	var newElement ListItem
	newElement.Value = v
	if l.len > 0 {
		newElement.Prev = l.last
		l.last.Next = &newElement
		l.last = &newElement
	} else {
		l.first, l.last = &newElement, &newElement
	}
	l.len++

	return l.last
}

func (l *list) Remove(i *ListItem) {
	if (l.len > 0) && (i != nil) {
		switch i {
		case l.first:
			l.first = i.Next
			l.first.Prev = nil
		case l.last:
			l.last = i.Prev
			l.last.Next = nil
		default:
			prev := i.Prev
			next := i.Next
			i.Prev.Next = next
			i.Next.Prev = prev
		}
		l.len--
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if l.len > 0 && i != nil {
		l.PushFront(i.Value)
		l.Remove(i)
	}
}

func NewList() List {
	return new(list)
}
