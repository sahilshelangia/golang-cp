package main

import (
	"fmt"

	"github.com/sahilshelangia/golang-cp/collections/stack"
)

func main() {
	stk := stack.New[int]()
	stk.Push(1)
	stk.Push(2)
	stk.Push(3)
	fmt.Println(stk.Size())
	fmt.Println(stk.Top())
}
