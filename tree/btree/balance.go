package btree

import "go-data-structure/constraints"

type Entry[K constraints.Ordered, V any] struct {
	key   K
	value V
}

func (e *Entry[K, V]) Key() K {
	return e.key
}

func (e *Entry[K, V]) Value() V {
	return e.value
}

type Node[K constraints.Ordered, V any] struct {
	entries  []*Entry[K, V]
	children []*Node[K, V]
}

func (n *Node[K, V]) isLeaf() bool {
	return len(n.children) == 0
}

func (n *Node[K, V]) writeAt(entry *Entry[K, V], off int) {
	n.entries = append(n.entries, nil)
	copy(n.entries[off+1:], n.entries[off:])
	n.entries[off] = entry
}

// find out index of entries where to append or continue to search children
func (n *Node[K, V]) binarySearch(key K) (int, bool) {
	left, right := 0, len(n.entries)-1
	for left <= right {
		mid := (left + right) / 2
		if key == n.entries[mid].key {
			return mid, true
		} else if key < n.entries[mid].key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left, false
}

type BTree[K constraints.Ordered, V any] struct {
	root   *Node[K, V]
	height int
	size   int
	m      int
}

func New[K constraints.Ordered, V any](m int) *BTree[K, V] {
	if m < 3 {
		panic("B-Tree: invalid M, should be at least 3")
	}
	return &BTree[K, V]{m: m}
}

func (t *BTree[K, V]) IsEmpty() bool {
	return t.size == 0
}

func (t *BTree[K, V]) Size() int {
	return t.size
}

func (t *BTree[K, V]) Height() int {
	return t.height
}

func (t *BTree[K, V]) Put(key K, value V) {
	if t.root == nil {
		t.root = &Node[K, V]{
			entries:  []*Entry[K, V](nil),
			children: []*Node[K, V](nil),
		}
		t.size++
		t.height++
		return
	}

	mid, right := t.put(t.root, &Entry[K, V]{key: key, value: value})
	if mid != nil {
		root := &Node[K, V]{
			entries:  []*Entry[K, V]{mid},
			children: []*Node[K, V]{t.root, right},
		}
		t.root = root
		t.height++
	}
}

func (t *BTree[K, V]) put(n *Node[K, V], entry *Entry[K, V]) (*Entry[K, V], *Node[K, V]) {
	index, ok := n.binarySearch(entry.key)
	if n.isLeaf() {
		// leaf node, insert entry
		if ok {
			n.entries[index] = entry
			return nil, nil
		}
		n.writeAt(entry, index)
		t.size++
	} else {
		// internal node, continue to find leaf node
		mid, right := t.put(n.children[index], entry)
		if mid == nil {
			return nil, nil
		}
		n.writeAt(mid, index)
		// backward shift && append 'right' child follows the 'index+1' position
		// keywords in new right child always greater than the 'index' position child
		n.children = append(n.children, nil)
		copy(n.children[index+1+1:], n.children[index+1:])
		n.children[index+1] = right
	}
	return t.split(n)
}

func (t *BTree[K, V]) split(n *Node[K, V]) (*Entry[K, V], *Node[K, V]) {
	if len(n.entries) < t.m {
		return nil, nil
	}
	// split left / middle / right parts
	mid := (t.m - 1) / 2
	right := &Node[K, V]{
		children: []*Node[K, V]{nil},
		entries:  append([]*Entry[K, V]{nil}, n.entries[mid+1:]...),
	}
	middle := n.entries[mid]
	n.entries = append([]*Entry[K, V]{nil}, n.entries[:mid]...)

	// if node is internal, split children also
	if !n.isLeaf() {
		n.children = append([]*Node[K, V](nil), n.children[:mid+1]...)
		right.children = append([]*Node[K, V](nil), n.children[mid+1:]...)
	}
	return middle, right
}
