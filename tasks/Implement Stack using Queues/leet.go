/*
Runtime
1 ms
Beats
62.89%

Memory
1.9 MB
Beats
57.73%
*/
package main

func main() {}

type MyStack struct {
	top    *node
	length int
}

type node struct {
	value int
	prev  *node
}

func Constructor() MyStack {
	return MyStack{nil, 0}
}

func (this *MyStack) Push(x int) {
	n := &node{x, this.top}
	this.top = n
	this.length++
}

func (this *MyStack) Pop() int {
	if this.length == 0 {
		return 0
	}

	n := this.top
	this.top = n.prev
	this.length--

	return n.value
}

func (this *MyStack) Top() int {
	if this.length == 0 {
		return 0
	}

	return this.top.value
}

func (this *MyStack) Empty() bool {
	if this.length == 0 {
		return true
	}

	return false
}

/**
* Your MyStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.Empty();
 */
