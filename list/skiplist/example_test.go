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

	for i := 0; i < 10; i++ {
		sl.Put(i, fn)
	}
	t.Log(sl)
}
