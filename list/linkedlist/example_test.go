package linkedlist

import "fmt"

func Output(l *LinkedList[int]) {
	l.ForEach(func(n *Node[int]) { fmt.Printf("%d -> ", n.Value) })
	fmt.Printf("%d", l.root.next.Value)
}

func LinkedList_New() *LinkedList[int] {
	l := New[int]()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	return l
}

func ExampleNew() {
	l := LinkedList_New()
	Output(l)
	// Output:
	// 0 -> 1 -> 2 -> 3 -> 4 -> 0
}

func ExampleLinkedList_Len() {
	l := LinkedList_New()
	fmt.Println(l.Len())
	// Output:
	// 5
}

func ExampleLinkedList_Front() {
	l := New[int]()
	fmt.Println(l.Front())

	l = LinkedList_New()
	fmt.Println(l.Front().Value)
	// Output:
	// <nil>
	// 0
}

func ExampleLinkedList_Back() {
	l := New[int]()
	fmt.Println(l.Back())

	l = LinkedList_New()
	fmt.Println(l.Back().Value)
	// Output:
	// <nil>
	// 4
}

func ExampleNode_Next() {
	l := LinkedList_New()
	fmt.Println(l.Front().Next().Value)
	// Output:
	// 1
}

func ExampleNode_Prev() {
	l := LinkedList_New()
	fmt.Println(l.Back().Prev().Value)
	// Output:
	// 3
}

func ExampleLinkedList_PushFront() {
	l := New[int]()
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}
	Output(l)
	// Output:
	// 4 -> 3 -> 2 -> 1 -> 0 -> 4
}

func ExampleLinkedList_InsertBefore() {
	l := LinkedList_New()
	n := l.Front().Next().Next()
	l.InsertBefore(8, n)
	Output(l)
	// Output:
	// 0 -> 1 -> 8 -> 2 -> 3 -> 4 -> 0
}

func ExampleLinkedList_InsertAfter() {
	l := LinkedList_New()
	n := l.Front().Next().Next()
	l.InsertAfter(8, n)
	Output(l)
	// Output:
	// 0 -> 1 -> 2 -> 8 -> 3 -> 4 -> 0
}

func ExampleLinkedList_MoveToFront() {
	l := LinkedList_New()
	n := l.Front().Next().Next()
	l.MoveToFront(n)
	Output(l)
	// Output:
	// 2 -> 0 -> 1 -> 3 -> 4 -> 2
}

func ExampleLinkedList_MoveToBack() {
	l := LinkedList_New()
	n := l.Front().Next().Next()
	l.MoveToBack(n)
	Output(l)
	// Output:
	// 0 -> 1 -> 3 -> 4 -> 2 -> 0
}

func ExampleLinkedList_MoveBefore() {
	l := LinkedList_New()
	n1 := l.Front().Next().Next()
	n2 := l.Front().Next()
	l.MoveBefore(n1, n2)
	Output(l)
	// Output:
	// 0 -> 2 -> 1 -> 3 -> 4 -> 0
}

func ExampleLinkedList_MoveAfter() {
	l := LinkedList_New()
	n1 := l.Front().Next().Next()
	n2 := l.Front().Next().Next().Next()
	l.MoveAfter(n1, n2)
	Output(l)
	// Output:
	// 0 -> 1 -> 3 -> 2 -> 4 -> 0
}

func ExampleLinkedList_Remove() {
	l := LinkedList_New()
	n := l.Front().Next().Next()
	l.Remove(n)
	Output(l)
	// Output:
	// 0 -> 1 -> 3 -> 4 -> 0
}

func ExampleLinkedList_ForEach() {
	l := LinkedList_New()
	l.ForEach(func(n *Node[int]) { n.Value = n.Value + 1 })
	Output(l)
	// Output:
	// 1 -> 2 -> 3 -> 4 -> 5 -> 1
}
