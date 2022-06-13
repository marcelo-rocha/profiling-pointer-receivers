package domain_test

import (
	"testing"
	"unsafe"

	"github.com/marcelo-rocha/profiling-pointer-receivers/domain"

	listbyref "github.com/marcelo-rocha/profiling-pointer-receivers/list/byref"
	listbyvalue "github.com/marcelo-rocha/profiling-pointer-receivers/list/byvalue"

	tuplebyref "github.com/marcelo-rocha/profiling-pointer-receivers/tuple/byref"
	tuplebyvalue "github.com/marcelo-rocha/profiling-pointer-receivers/tuple/byvalue"
)

func TestListByReference(t *testing.T) {
	l := listbyvalue.LeafNode{}
	r1 := unsafe.Sizeof(l)

	tuple := tuplebyvalue.LeafNode{}
	r2 := unsafe.Sizeof(tuple)

	t.Log("Size:", r1, r2)
}

func BenchmarkListByReferenceL5(b *testing.B) {
	root := domain.MakeCollection5(listbyref.NewInnerNode, listbyref.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkListByValueL5(b *testing.B) {
	root := domain.MakeCollection5(listbyvalue.NewInnerNode, listbyvalue.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkTupleByReferenceL5(b *testing.B) {
	root := domain.MakeCollection5(tuplebyref.NewInnerNode, tuplebyref.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkTupleByValueL5(b *testing.B) {
	root := domain.MakeCollection5(tuplebyvalue.NewInnerNode, tuplebyvalue.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkListByReferenceL3(b *testing.B) {
	root := domain.MakeCollection3(listbyref.NewInnerNode, listbyref.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkListByValueL3(b *testing.B) {
	root := domain.MakeCollection3(listbyvalue.NewInnerNode, listbyvalue.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkTupleByReferenceL3(b *testing.B) {
	root := domain.MakeCollection3(tuplebyref.NewInnerNode, tuplebyref.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkTupleByValueL3(b *testing.B) {
	root := domain.MakeCollection3(tuplebyvalue.NewInnerNode, tuplebyvalue.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkListByReferenceL2(b *testing.B) {
	root := domain.MakeCollection2(listbyref.NewInnerNode, listbyref.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkListByValueL2(b *testing.B) {
	root := domain.MakeCollection2(listbyvalue.NewInnerNode, listbyvalue.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkTupleByReferenceL2(b *testing.B) {
	root := domain.MakeCollection2(tuplebyref.NewInnerNode, tuplebyref.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkTupleByValueL2(b *testing.B) {
	root := domain.MakeCollection2(tuplebyvalue.NewInnerNode, tuplebyvalue.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}
