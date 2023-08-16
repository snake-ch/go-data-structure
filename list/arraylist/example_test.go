package arraylist

import "fmt"

func Output(l *ArrayList[int]) {
	fmt.Println(l.elements)
}

func ExampleArrayList_Len() {
	l := New(1, 2, 3, 4, 5)
	Output(l)
	fmt.Println(l.Len())
	// Output:
	// [1 2 3 4 5]
	// 5
}

func ExampleArrayList_Cap() {
	l := New(1, 2, 3, 4, 5)
	Output(l)
	fmt.Println(l.Cap())
	// Output:
	// [1 2 3 4 5]
	// 32
}

func ExampleArrayList_Append() {
	l := New(1, 2, 3, 4, 5)
	l.Append(9, 8, 7, 6, 5)
	Output(l)
	// Output:
	// [1 2 3 4 5 9 8 7 6 5]
}

func ExampleArrayList_Get() {
	l := New(1, 2, 3, 4, 5)
	for idx := 0; idx < 5; idx++ {
		e, _ := l.Get(idx)
		fmt.Println(e)
	}
	e, ok := l.Get(-1)
	fmt.Println(e, ok)
	e, ok = l.Get(l.Len())
	fmt.Println(e, ok)
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 0 false
	// 0 false
}

func ExampleArrayList_Set() {
	l := New(1, 2, 3, 4, 5)
	l.Set(-1, 0)
	l.Set(l.Len(), 0)
	for i := 0; i < 5; i++ {
		l.Set(i, i+2)
	}
	Output(l)
	// Output:
	// [2 3 4 5 6]
}

func ExampleArrayList_Remove() {
	l := New(1, 2, 3, 4, 5)
	l.Remove(-1)
	l.Remove(l.Len())
	l.Remove(2)
	Output(l)
	// Output:
	// [1 2 4 5]
}

func ExampleArrayList_Foreach() {
	l := New(1, 2, 3, 4, 5)
	l.Foreach(func(e int) {
		fmt.Printf("%d ", e*e)
	})
	// Output:
	// 1 4 9 16 25
}

func ExampleArrayList_Contains() {
	l := New(1, 2, 3, 4, 5)
	idx0, exist0 := l.Contains(func(i, j int) bool { return i == j }, 2)
	idx1, exist1 := l.Contains(func(i, j int) bool { return i == j }, 10)
	fmt.Println(idx0, exist0)
	fmt.Println(idx1, exist1)
	// Output:
	// 1 true
	// -1 false
}
