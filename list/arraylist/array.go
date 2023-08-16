package arraylist

const DEFAULT_CAP = 32

type ArrayList[T any] struct {
	elements []T
}

func New[T any](e ...T) *ArrayList[T] {
	al := &ArrayList[T]{
		elements: make([]T, 0, DEFAULT_CAP),
	}
	al.elements = append(al.elements, e...)
	return al
}

func (l *ArrayList[T]) Len() int {
	return len(l.elements)
}

func (l *ArrayList[T]) Cap() int {
	return cap(l.elements)
}

func (l *ArrayList[T]) Append(e ...T) {
	l.elements = append(l.elements, e...)
}

func (l *ArrayList[T]) Insert(idx int, e ...T) {
	if idx < 0 || idx >= l.Len() {
		l.elements = append(l.elements, e...)
	} else {
		l.elements = append(l.elements[:idx], append(e, l.elements[idx:]...)...)
	}
}

// return empty of 'T' if out of range list orelse the replaced one
func (l *ArrayList[T]) Set(idx int, e T) (T, bool) {
	var el T
	if idx < 0 || idx >= l.Len() {
		return el, false
	}
	el = l.elements[idx]
	l.elements[idx] = e
	return el, true
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

func (l *ArrayList[T]) Foreach(fn func(e T)) {
	for _, e := range l.elements {
		fn(e)
	}
}

// returns first position index of element and if present or not
func (l *ArrayList[T]) Contains(compare func(i, j T) bool, es ...T) (int, bool) {
	for _, e := range es {
		for idx, el := range l.elements {
			if compare(el, e) {
				return idx, true
			}
		}
	}
	return -1, false
}
