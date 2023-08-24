package avltree

import (
	"testing"
)

func TestAvltree(t *testing.T) {
	tree := New[string, int]()
	for char := 'a'; char <= 'z'; char++ {
		tree.Put(string(char), 0)
	}
	t.Log(tree)
}
