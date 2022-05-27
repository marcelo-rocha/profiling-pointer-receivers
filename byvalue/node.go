package byvalue

import "github.com/marcelo-rocha/profiling-pointer-receivers/common"

type InnerNode struct {
	child      common.Mapper
	values     common.ValuesArray
	loopLength int
}

func NewInnerNode(values common.ValuesArray, child common.Mapper, loopLength int) *InnerNode {
	return &InnerNode{
		child:      child,
		values:     values,
		loopLength: loopLength,
	}
}

func (n InnerNode) Add(total []int64) {
	for i := 0; i < n.loopLength; i++ {
		n.child.Add(total)
	}
	for i := range n.values {
		total[i] += n.values[i]
	}
}

type LeafNode struct {
	values common.ValuesArray
}

func NewLeafNode(values common.ValuesArray) *LeafNode {
	return &LeafNode{
		values: values,
	}
}

func (n LeafNode) Add(total []int64) {
	for i := range n.values {
		total[i] += n.values[i]
	}
}
