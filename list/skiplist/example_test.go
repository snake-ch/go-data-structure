package skiplist

import "testing"

func TestSkiplist(t *testing.T) {
	sl := New[int]()
	fn := func(i, j int) int {
		if i < j {
			return -1
		}
		if i == j {
			return 0
		}
		return 1
	}

	sl.Put(4, fn)
	sl.Put(8, fn)
	sl.Put(6, fn)
	sl.Put(5, fn)
	sl.Put(1, fn)
	t.Log(sl)
}
