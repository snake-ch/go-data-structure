package linkedlist

// Node is an element of a linked list.
type Node[T any] struct {
	prev, next *Node[T]
	Value      T
}

func (n *Node[T]) Prev() *Node[T] {
	if n.prev == n {
		return nil
	}
	return n.prev
}

func (n *Node[T]) Next() *Node[T] {
	if n.next == n {
		return nil
	}
	return n.next
}

// LinkedList represents a doubly linked list.
type LinkedList[T any] struct {
	root   Node[T]
	length int
}

func New[T any]() *LinkedList[T] {
	l := &LinkedList[T]{}
	l.root.prev = &l.root
	l.root.next = &l.root
	return l
}

func (l *LinkedList[T]) Len() int { return l.length }

func (l *LinkedList[T]) Front() *Node[T] {
	if l.length == 0 {
		return nil
	}
	return l.root.next
}

func (l *LinkedList[T]) Back() *Node[T] {
	if l.length == 0 {
		return nil
	}
	return l.root.prev
}

func (l *LinkedList[T]) insert(v T, n *Node[T]) {
	e := &Node[T]{
		prev:  n,
		next:  n.next,
		Value: v,
	}
	e.prev.next = e
	e.next.prev = e
	l.length = l.length + 1
}

func (l *LinkedList[T]) move(e, n *Node[T]) {
	if e == nil || e == n {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	e.prev = n
	e.next = n.next
	e.prev.next = e
	e.next.prev = e
}

func (l *LinkedList[T]) PushFront(v T)                { l.insert(v, &l.root) }
func (l *LinkedList[T]) PushBack(v T)                 { l.insert(v, l.root.prev) }
func (l *LinkedList[T]) InsertBefore(v T, n *Node[T]) { l.insert(v, n.prev) }
func (l *LinkedList[T]) InsertAfter(v T, n *Node[T])  { l.insert(v, n) }
func (l *LinkedList[T]) MoveToFront(n *Node[T])       { l.move(n, &l.root) }
func (l *LinkedList[T]) MoveToBack(n *Node[T])        { l.move(n, l.root.prev) }
func (l *LinkedList[T]) MoveBefore(e, n *Node[T])     { l.move(e, n.prev) }
func (l *LinkedList[T]) MoveAfter(e, n *Node[T])      { l.move(e, n) }

func (l *LinkedList[T]) Remove(n *Node[T]) {
	if n == nil {
		return
	}
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil
	l.length = l.length - 1
}

func (l *LinkedList[T]) ForEach(fn func(*Node[T])) {
	for n := l.Front(); n != &l.root; n = n.Next() {
		fn(n)
	}
}
