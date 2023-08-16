package arraylist

type ArrayList[T any] struct {
	elements []T
}

func New[T any]() *ArrayList[T] {
	return &ArrayList[T]{
		elements: make([]T, 0, 32),
	}
}

func (l *ArrayList[T]) Len() int {
	return len(l.elements)
}

func (l *ArrayList[T]) Add(idx int, e T) {
	if idx < 0 || idx > l.Len() {
		return
	}
	if idx == 0 {
		l.elements = append([]T{e}, l.elements...)
		return
	}
	if idx == l.Len() {
		l.elements = append(l.elements, e)
		return
	}
	l.elements = append(l.elements[:idx], append([]T{e}, l.elements[idx:]...)...)
}

func (l *ArrayList[T]) Set(idx int, e T) T {
	var el T
	if idx < 0 || idx >= l.Len() {
		return el
	}
	el = l.elements[idx]
	l.elements[idx] = e
	return el
}

func (l *ArrayList[T]) Get(idx int) (T, bool) {
	var e T
	if idx < 0 || idx >= l.Len() {
		return e, false
	}
	return l.elements[idx], true
}

func (l *ArrayList[T]) Remove(idx int) {
	if idx < 0 || idx >= l.Len() {
		return
	}
	l.elements = l.elements[:idx+copy(l.elements[idx:], l.elements[idx+1:])]
}

// returns position index of element and if present or not
func (l *ArrayList[T]) Contains(e T, fn func(i, j T) bool) (int, bool) {
	for idx, el := range l.elements {
		if fn(e, el) {
			return idx, true
		}
	}
	return -1, false
}
