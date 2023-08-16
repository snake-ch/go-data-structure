package skiplist

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const _MAX_LEVEL int = 32
const _FACTOR float64 = 0.5

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}

type Node[T any] struct {
	Value    T
	backward *Node[T]
	level    []struct {
		forward *Node[T]
		span    int
	}
}

type SkipList[T any] struct {
	header, tail *Node[T]
	length       int
	level        int
}

func New[T any]() *SkipList[T] {
	header := &Node[T]{
		Value:    *new(T),
		backward: nil,
		level: make([]struct {
			forward *Node[T]
			span    int
		}, _MAX_LEVEL),
	}

	return &SkipList[T]{
		header: header,
		tail:   nil,
		length: 0,
		level:  1,
	}
}

func randomLevel() int {
	level := 1
	for r.Float64() < _FACTOR && level < _MAX_LEVEL {
		level = level + 1
	}
	return level
}

func createNode[T any](v T) *Node[T] {
	return &Node[T]{
		Value:    v,
		backward: nil,
		level: make([]struct {
			forward *Node[T]
			span    int
		}, randomLevel()),
	}
}

func (sl *SkipList[T]) Put(v T, compare func(i, j T) int) {
	prev := make([]*Node[T], _MAX_LEVEL)
	rank := make([]int, _MAX_LEVEL)

	// predecessors less than 'v' in each level where to insert & store rank
	n := sl.header
	for l := sl.level - 1; l >= 0; l-- {
		if l == sl.level-1 {
			rank[l] = 0
		} else {
			rank[l] = rank[l+1]
		}
		for n.level[l].forward != nil && compare(n.level[l].forward.Value, v) < 0 {
			rank[l] += n.level[l].span
			n = n.level[l].forward
		}
		prev[l] = n
	}

	n = createNode(v)
	level := len(n.level)
	if sl.level < level {
		for l := sl.level; l < level; l++ {
			rank[l] = 0
			prev[l] = sl.header
			prev[l].level[l].span = sl.length
		}
		sl.level = level
	}
	for l := 0; l < level; l++ {
		// insert node
		n.level[l].forward = prev[l].level[l].forward
		prev[l].level[l].forward = n
		// update span
		n.level[l].span = prev[l].level[l].span - (rank[0] - rank[l])
		prev[l].level[l].span = rank[0] - rank[l] + 1
	}
	// increase span
	for l := level; l < sl.level; l++ {
		prev[l].level[l].span++
	}
	// update relation
	if prev[0] == sl.header {
		n.backward = nil
	} else {
		n.backward = prev[0]
	}
	if n.level[0].forward != nil {
		n.level[0].forward.backward = n
	} else {
		sl.tail = n
	}
	sl.length++
}

func (sl *SkipList[T]) Get(v T, compare func(i, j T) int) *Node[T] {
	n := sl.header
	for l := sl.level - 1; l >= 0; l-- {
		for n.level[l].forward != nil && compare(v, n.level[l].forward.Value) < 0 {
			n = n.level[l].forward
		}
	}
	return n.level[0].forward
}

func (sl *SkipList[T]) Remove(v T, compare func(i, j T) int) {
	prev := make([]*Node[T], _MAX_LEVEL)

	// find all predecessors of 'v'
	n := sl.header
	for l := sl.level - 1; l >= 0; l-- {
		for n.level[l].forward != nil && compare(v, n.level[l].forward.Value) < 0 {
			n = n.level[l].forward
		}
		prev[l] = n
	}

	// node with value 'v' is presented or not
	n = n.level[0].forward
	if n == nil || compare(v, n.Value) != 0 {
		return
	}

	// update predecessors
	for l := 0; l < sl.level; l++ {
		if prev[l].level[l].forward == n {
			prev[l].level[l].span += n.level[l].span - 1
			prev[l].level[l].forward = n.level[l].forward
		} else {
			prev[l].level[l].span--
		}
	}
	if n.level[0].forward != nil {
		n.level[0].forward.backward = n.backward
	} else {
		sl.tail = n.backward
	}
	// if remove top level node
	if sl.level > 1 && sl.header.level[sl.level-1].forward == nil {
		sl.level--
	}
	sl.length--
}

func (sl *SkipList[T]) String() string {
	sb := new(strings.Builder)
	sb.WriteString("\n\n")
	for l := sl.level - 1; l >= 0; l-- {
		sb.WriteString("lv")
		sb.WriteString(fmt.Sprintf("%-2v", l))
		sb.WriteString(fmt.Sprintf("%-2s", ":"))

		var e *Node[T]
		for e = sl.header; e != nil; e = e.level[l].forward {
			sb.WriteString(fmt.Sprintf("%-3v", e.Value))
			for i := 0; i < e.level[l].span; i++ {
				if i == 0 {
					sb.WriteString(fmt.Sprintf("%-3s", "~"))
				} else {
					sb.WriteString(fmt.Sprintf("%-3s%-3s", "", "~"))
				}
			}
			if e.level[l].forward == nil && e.level[l].span != 0 {
				sb.WriteString(fmt.Sprintf("%-3s", ""))
			}
		}
		sb.WriteString("--> nil\n")
	}
	return sb.String()
}
