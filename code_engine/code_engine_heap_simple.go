package code_engine

import "math"

type OutHeap interface {
	Add(node *Node)
	GetMin() *Node
}

// OutHeapSimple is min heap by default
type OutHeapSimple struct {
	data []*Node
}

func (oh *OutHeapSimple) Add(node *Node) {
	oh.data = append(oh.data, node)
}

func (oh *OutHeapSimple) GetMin() *Node {
	if len(oh.data) != 0 {
		// search min node
		var minIndex int
		minValue := math.MaxInt64
		for index, node := range oh.data {
			if node.val < minValue {
				minValue = node.val
				minIndex = index
			}
		}
		d := oh.data
		// swap min node and last node, simulate heap operation
		d[len(d)-1], d[minIndex] = d[minIndex], d[len(d)-1]
		// truncate last node and return it
		node := d[len(d)-1]
		oh.data = d[:len(d)-1]
		return node
	}
	return nil
}
