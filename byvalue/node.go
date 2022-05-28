package byvalue

import "github.com/marcelo-rocha/profiling-pointer-receivers/tuple"

type InnerNode struct {
	child         tuple.Mapper
	tuple         tuple.Tuple
	iterationsQty int
}

func NewInnerNode(values tuple.Tuple, child tuple.Mapper, iterationsQty int) *InnerNode {
	return &InnerNode{
		child:         child,
		tuple:         values,
		iterationsQty: iterationsQty,
	}
}

func (n InnerNode) Add(total []int64) {
	for i := 0; i < n.iterationsQty; i++ {
		n.child.Add(total)
	}
	for i := range n.tuple {
		total[i] += n.tuple[i]
	}
}

type LeafNode struct {
	tuple tuple.Tuple
}

func NewLeafNode(values tuple.Tuple) *LeafNode {
	return &LeafNode{
		tuple: values,
	}
}

func (n LeafNode) Add(total []int64) {
	for i := range n.tuple {
		total[i] += n.tuple[i]
	}
}
