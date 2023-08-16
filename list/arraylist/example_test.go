package arraylist

import "fmt"

func Output(l *ArrayList[int]) {
	fmt.Println(l.elements)
}

func ArrayList_New() *ArrayList[int] {
	l := New[int]()
	for i := 0; i < 5; i++ {
		l.Add(l.Len(), i)
	}
	return l
}

func ExampleNew() {
	l := ArrayList_New()
	Output(l)
	// Output:
	// [0 1 2 3 4]
}

func ExampleArrayList_Len() {
	l := ArrayList_New()
	fmt.Println(l.Len())
	// Output:
	// 5
}

func ExampleArrayList_Add() {
	l := ArrayList_New()
	l.Add(-1, 0)
	l.Add(l.Len()+1, 0)
	l.Add(2, 0)
	Output(l)
	// Output:
	// [0 1 0 2 3 4]
}

func ExampleArrayList_Get() {
	l := ArrayList_New()
	for idx := 0; idx < 5; idx++ {
		e, _ := l.Get(idx)
		fmt.Println(e)
	}
	e, ok := l.Get(-1)
	fmt.Println(e, ok)
	e, ok = l.Get(l.Len())
	fmt.Println(e, ok)
	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
	// 0 false
	// 0 false
}

func ExampleArrayList_Set() {
	l := ArrayList_New()
	l.Set(-1, 0)
	l.Set(l.Len(), 0)
	for i := 0; i < 5; i++ {
		l.Set(i, i+1)
	}
	Output(l)
	// Output:
	// [1 2 3 4 5]
}

func ExampleArrayList_Remove() {
	l := ArrayList_New()
	l.Remove(-1)
	l.Remove(l.Len())
	l.Remove(2)
	Output(l)
	// Output:
	// [0 1 3 4]
}

func ExampleArrayList_Contains() {
	l := ArrayList_New()
	idx0, exist0 := l.Contains(2, func(i, j int) bool { return i == j })
	idx1, exist1 := l.Contains(10, func(i, j int) bool { return i == j })
	fmt.Println(idx0, exist0)
	fmt.Println(idx1, exist1)
	// Output:
	// 2 true
	// -1 false
}
