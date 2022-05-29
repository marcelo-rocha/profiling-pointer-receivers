package domain_test

import (
	"testing"

	"github.com/marcelo-rocha/profiling-pointer-receivers/domain"

	listbyref "github.com/marcelo-rocha/profiling-pointer-receivers/list/byref"
	listbyvalue "github.com/marcelo-rocha/profiling-pointer-receivers/list/byvalue"

	tuplebyref "github.com/marcelo-rocha/profiling-pointer-receivers/tuple/byref"
	tuplebyvalue "github.com/marcelo-rocha/profiling-pointer-receivers/tuple/byvalue"
)

func BenchmarkListByReference(b *testing.B) {
	root := domain.MakeCollection(listbyref.NewInnerNode, listbyref.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkListByValue(b *testing.B) {
	root := domain.MakeCollection(listbyvalue.NewInnerNode, listbyvalue.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkTupleByReference(b *testing.B) {
	root := domain.MakeCollection(tuplebyref.NewInnerNode, tuplebyref.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}

func BenchmarkTupleByValue(b *testing.B) {
	root := domain.MakeCollection(tuplebyvalue.NewInnerNode, tuplebyvalue.NewLeafNode)
	values := make([]int64, domain.ValuesQty)
	root.Add(values)
}
