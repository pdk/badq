package main

import (
	"fmt"

	"github.com/pdk/badq/queue"
)

func main() {

	q := queue.New[int]()

	q.Push(1)
	q.Push(2)
	fmt.Printf("%d\n", q.Pull())
	q.Push(3)
	q.Push(4)
	q.Push(5)
	fmt.Printf("%d\n", q.Pull())
	fmt.Printf("%d\n", q.Pull())
	q.Push(6)
	fmt.Printf("%d\n", q.Pull())
	q.Push(7)
	q.Push(8)
	fmt.Printf("%d\n", q.Pull())
	q.Push(9)
	q.Push(10)

	for !q.Empty() {
		fmt.Printf("%d\n", q.Pull())
	}

	fmt.Printf("%d\n", q.Pull())
}
