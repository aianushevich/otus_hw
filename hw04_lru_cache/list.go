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
	slice []*ListItem
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
	l.slice = append(l.slice, &newElement)
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
	l.slice = append(l.slice, &newElement)
	l.len++

	return l.last
}

func (l *list) Remove(i *ListItem) {
	if (l.len > 0) && (i != nil) {
		switch i {
		case l.first:
			{
				l.first = i
				l.first.Prev = nil
				l.len--
			}
		case l.last:
			{
				l.last = i
				l.last.Next = nil
				l.len--
			}
		default:
			{
				prev := i.Prev
				next := i.Next
				i.Prev.Next = next
				i.Next.Prev = prev
				l.len--
			}
		}
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if (l.len > 1) && (i != l.first) && (i != l.last) {
		prev := i.Prev
		next := i.Next

		i.Prev.Next = next
		i.Next.Prev = prev

		i.Prev = nil
		i.Next = l.first
		l.first.Prev = i
	} else if (l.len > 1) && (i == l.last) {
		i.Prev.Next = nil
		i.Prev = nil
		i.Next = l.first
		l.first.Prev = i
		l.first = i
	}
}

func NewList() List {
	return new(list)
}
