package byref

import (
	"github.com/marcelo-rocha/profiling-pointer-receivers/domain"
	"github.com/marcelo-rocha/profiling-pointer-receivers/list"
)

type InnerNode struct {
	child         domain.Mapper
	list          list.List
	iterationsQty int
}

func NewInnerNode(values []int64, child domain.Mapper, iterationsQty int) domain.Mapper {
	r := &InnerNode{
		child:         child,
		iterationsQty: iterationsQty,
	}
	r.list[0] = values[0]
	r.list[1] = values[1]
	r.list[2] = values[2]
	r.list[3] = values[3]
	r.list[4] = values[4]
	return r
}

func (n *InnerNode) Add(total []int64) {
	for i := 0; i < n.iterationsQty; i++ {
		n.child.Add(total)
	}
	total[0] += n.list[0]
	total[1] += n.list[1]
	total[2] += n.list[2]
	total[3] += n.list[3]
	total[4] += n.list[4]
}

type LeafNode struct {
	list list.List
}

func NewLeafNode(values []int64) domain.Mapper {
	r := &LeafNode{}
	r.list[0] = values[0]
	r.list[1] = values[1]
	r.list[2] = values[2]
	r.list[3] = values[3]
	r.list[4] = values[4]
	return r
}

func (n *LeafNode) Add(total []int64) {
	total[0] += n.list[0]
	total[1] += n.list[1]
	total[2] += n.list[2]
	total[3] += n.list[3]
	total[4] += n.list[4]
}
