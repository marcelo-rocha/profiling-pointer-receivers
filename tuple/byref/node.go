package byref

import (
	"github.com/marcelo-rocha/profiling-pointer-receivers/domain"
	"github.com/marcelo-rocha/profiling-pointer-receivers/tuple"
)

type InnerNode struct {
	tuple.Tuple
	child         domain.Mapper
	iterationsQty int
}

func NewInnerNode(values []int64, child domain.Mapper, iterationsQty int) domain.Mapper {
	r := &InnerNode{
		child:         child,
		iterationsQty: iterationsQty,
	}
	r.Tuple.V0 = values[0]
	r.Tuple.V1 = values[1]
	r.Tuple.V2 = values[2]
	r.Tuple.V3 = values[3]
	r.Tuple.V4 = values[4]
	return r
}

func (n *InnerNode) Add(total []int64) {
	for i := 0; i < n.iterationsQty; i++ {
		n.child.Add(total)
	}
	total[0] += n.Tuple.V0
	total[1] += n.Tuple.V1
	total[2] += n.Tuple.V2
	total[3] += n.Tuple.V3
	total[4] += n.Tuple.V4
}

type LeafNode struct {
	tuple.Tuple
}

func NewLeafNode(values []int64) domain.Mapper {
	r := &LeafNode{}
	r.Tuple.V0 = values[0]
	r.Tuple.V1 = values[1]
	r.Tuple.V2 = values[2]
	r.Tuple.V3 = values[3]
	r.Tuple.V4 = values[4]
	return r
}

func (n *LeafNode) Add(total []int64) {
	total[0] += n.Tuple.V0
	total[1] += n.Tuple.V1
	total[2] += n.Tuple.V2
	total[3] += n.Tuple.V3
	total[4] += n.Tuple.V4
}
