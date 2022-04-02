/*
Runtime: 42 ms, faster than 22.90% of Go online submissions for Min Stack.
Memory Usage: 9.1 MB, less than 19.37% of Go online submissions for Min Stack.
*/
package main

import (
	"container/list"
	"fmt"
)

func main() {
	t := Constructor()
	fmt.Println(t)
	t.Push(-2)
	t.Push(0)
	t.Push(-3)
	fmt.Println(t)
	fmt.Println(t.GetMin())
	t.Pop()
	fmt.Println(t.Top())
	fmt.Println(t.GetMin())
}

type pair struct {
	val int
	min int
}

type MinStack struct {
	stack list.List
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	newPair := pair{val: val, min: val}

	if this.stack.Len() > 0 {
		lastMin := this.GetMin()
		if lastMin < val {
			newPair.min = lastMin
		}
	}
	this.stack.PushBack(newPair)
}

func (this *MinStack) Pop() {
	this.stack.Remove(this.stack.Back())
}

func (this *MinStack) Top() int {
	return this.stack.Back().Value.(pair).val
}

func (this *MinStack) GetMin() int {
	return this.stack.Back().Value.(pair).min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
