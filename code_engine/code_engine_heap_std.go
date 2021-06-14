package code_engine

import "container/heap"

// OutHeapStd is min heap by default
type OutHeapStd struct {
	data []*Node
}

func (oh *OutHeapStd) Len() int {
	return len(oh.data)
}

func (oh *OutHeapStd) Less(i, j int) bool {
	return oh.data[i].val < oh.data[j].val
}

func (oh *OutHeapStd) Swap(i, j int) {
	oh.data[i], oh.data[j] = oh.data[j], oh.data[i]
}

func (oh *OutHeapStd) Push(x interface{}) {
	oh.data = append(oh.data, x.(*Node))
}

func (oh *OutHeapStd) Pop() interface{} {
	d := oh.data
	x := d[len(d)-1]
	oh.data = d[:len(d)-1]
	d[len(d)-1] = nil
	return x
}

// Add use std library heap operations
func (oh *OutHeapStd) Add(node *Node) {
	heap.Push(oh, node)
}

func (oh *OutHeapStd) GetMin() *Node {
	x := heap.Pop(oh)
	node := x.(*Node)
	return node
}
