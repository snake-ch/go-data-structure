package avltree

import (
	"fmt"
	"math"
	"strings"

	"go-data-structure/constraints"
)

const (
	LEFT = iota
	RIGHT
)

type AvlTree[K constraints.Ordered, V any] struct {
	root *Node[K, V]
}

func New[K constraints.Ordered, V any]() *AvlTree[K, V] {
	return &AvlTree[K, V]{root: nil}
}

func (t *AvlTree[K, V]) Put(key K, value V) {
	t.root = t.root.put(key, value, t.root)
}

func (t *AvlTree[K, V]) Get(key K) (value V, exist bool) {
	n := t.root.get(key)
	if n == nil {
		return
	}
	return n.value, true
}

func (t *AvlTree[K, V]) Remove(key K) {
	t.root = t.root.remove(key)
}

func (t *AvlTree[K, V]) Height() int {
	return t.root.height()
}

func (t *AvlTree[K, V]) Size() int {
	return t.root.size()
}

func (t *AvlTree[K, V]) String() string {
	height := t.Height()
	width := int(6*math.Pow(2, float64(height-2)) - 1)
	array := make([][]string, height*2)
	for row := range array {
		array[row] = make([]string, width)
		for col := range array[row] {
			array[row][col] = " "
		}
	}
	print(t.root, 0, (width-1)/2, height, array)

	sb := new(strings.Builder)
	sb.WriteString("\n\n")
	for row := range array {
		for col := range array[row] {
			sb.WriteString(fmt.Sprintf("%v", array[row][col]))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func print[K constraints.Ordered, V any](n *Node[K, V], row, col, height int, array [][]string) {
	if n == nil {
		return
	}
	array[row][col] = fmt.Sprintf("%v", n.key)

	if row/2+1 == height {
		return
	}

	gap := 0
	if height-2-row/2 != 0 {
		width := int(6*math.Pow(2, float64(height-3-row/2)) - 1)
		gap = (width - 1) / 2
	} else {
		gap = 1
	}

	if n.left != nil {
		array[row+1][col-(gap-1)/2-1] = "/"
		print(n.left, row+2, col-gap-1, height, array)
	}
	if n.right != nil {
		array[row+1][col+(gap+1)/2] = "\\"
		print(n.right, row+2, col+gap+1, height, array)
	}
}

type Node[K constraints.Ordered, V any] struct {
	key         K
	value       V
	h           int
	parent      *Node[K, V]
	left, right *Node[K, V]
}

func (n *Node[K, V]) put(key K, value V, parent *Node[K, V]) *Node[K, V] {
	if n == nil {
		return &Node[K, V]{
			key:    key,
			value:  value,
			h:      1,
			parent: parent,
			left:   nil,
			right:  nil,
		}
	}
	if key < n.key {
		n.left = n.left.put(key, value, n)
	} else if key > n.key {
		n.right = n.right.put(key, value, n)
	} else {
		n.value = value
	}
	return n.rebalance()
}

func (n *Node[K, V]) get(key K) *Node[K, V] {
	if n == nil {
		return nil
	}
	if key < n.key {
		return n.left.get(key)
	} else if key > n.key {
		return n.right.get(key)
	}
	return n

}

func (n *Node[K, V]) remove(key K) *Node[K, V] {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = n.left.remove(key)
	} else if key > n.key {
		n.right = n.right.remove(key)
	} else {
		if n.left == nil && n.right == nil { // case 1: no child
			n = nil
		} else if n.right == nil { // case 2: left child only
			n = n.left
		} else if n.left == nil { // case 3: right child only
			n = n.right
		} else {
			// Case 4: both left and right child, right child is not a leaf
			//   Step 1. find the node N with the smallest key
			//           and its parent P on the right subtree
			//   Step 2. swap S and N
			//   Step 3. remove node N like Case 1 or Case 3
			//   Step 4. update height for P
			//     |                  |
			//     N                  S                 |
			//    / \                / \                S
			//   L  ..  swap(N, S)  L  ..  remove(N)   / \
			//       |  =========>      |  ========>  L  ..
			//       P                  P                 |
			//      / \                / \                P
			//     S  ..              N  ..              / \
			//      \                  \                R  ..
			//       R                  R
			successor := n.right
			for successor.left != nil {
				successor = successor.left
			}
			n.key = successor.key
			n.value = successor.value
			n.right = n.right.remove(successor.key)
		}
	}
	return n.rebalance()
}

func (n *Node[K, V]) rebalance() *Node[K, V] {
	if n == nil {
		return nil
	}

	n.adjustHeight()
	unbalance := n.left.height() - n.right.height()
	if unbalance > 1 {
		// LR: transform to LL by rotate right-child left
		//     |                   |
		//     C                   C                 |
		//    /   l-rotate(A)     /   r-rotate(C)    B
		//   A    ==========>    B    ==========>   / \
		//    \                 /                  A   C
		//     B               A
		if n.left.right.height() > n.left.left.height() {
			n.left = n.left.rotate(LEFT)
		}
		// LL: fixed by rotate right
		//       |
		//       C                 |
		//      /   r-rotate(C)    B
		//     B    ==========>   / \
		//    /                  A   C
		//   A
		return n.rotate(RIGHT)
	} else if unbalance < -1 {
		// RL: transform to RR by rotate left-child right
		//   |                 |
		//   A                 A                     |
		//    \   r-rotate(C)   \     l-rotate(A)    B
		//     C  ==========>    B    ==========>   / \
		//    /                   \                A   C
		//   B                     C
		if n.right.left.height() > n.right.right.height() {
			n.right = n.right.rotate(RIGHT)
		}
		// RR: fixed by rotate left
		//   |
		//   C                     |
		//    \     l-rotate(C)    B
		//     B    ==========>   / \
		//      \                A   C
		//       A
		return n.rotate(LEFT)
	}
	return n
}

func (n *Node[K, V]) height() int {
	if n == nil {
		return 0
	}
	return n.h
}

func (n *Node[K, V]) adjustHeight() {
	n.h = 1 + int(math.Max(float64(n.left.height()), float64(n.right.height())))
}

// rotate left:
//
//	   G                     P
//	 /   \                 /   \
//	?     P               G     ?
//	     / \     -->     / \     \
//	    ?   ?           ?   ?     ?
//	         \
//	          ?
//
// rotate right:
//
//	       G                   P
//	     /   \               /   \
//	    P     ?             ?     G
//	   / \         -->     /     / \
//	  ?   ?               ?     ?   ?
//	 /
//	?
func (n *Node[K, V]) rotate(lr int) (root *Node[K, V]) {
	switch lr {
	case LEFT:
		root = n.right
		n.right = root.left
		root.left = n

		// update parent
		root.parent = n.parent
		n.parent = root
		if n.right != nil {
			n.right.parent = n
		}
	case RIGHT:
		root = n.left
		n.left = root.right
		root.right = n

		// update parent
		root.parent = n.parent
		n.parent = root
		if n.left != nil {
			n.left.parent = n
		}
	}
	n.adjustHeight()
	root.adjustHeight()
	return
}

func (n *Node[K, V]) size() int {
	if n == nil {
		return 0
	}
	return 1 + n.left.size() + n.right.size()
}

func (n *Node[K, V]) String() string {
	return fmt.Sprintf("%v", n.key)
}
