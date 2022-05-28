package byvalue_test

import (
	"testing"

	"github.com/marcelo-rocha/profiling-pointer-receivers/byvalue"
	"github.com/marcelo-rocha/profiling-pointer-receivers/params"
	"github.com/marcelo-rocha/profiling-pointer-receivers/tuple"
)

func BenchmarkByValue(b *testing.B) {
	initialValues := tuple.TupleOnes
	var root tuple.Mapper = byvalue.NewInnerNode(initialValues,
		byvalue.NewInnerNode(initialValues,
			byvalue.NewInnerNode(initialValues,
				byvalue.NewInnerNode(initialValues,
					byvalue.NewLeafNode(initialValues),
					params.InnerIterationsQty),
				params.InnerIterationsQty),
			params.InnerIterationsQty),
		params.RootIterationsQty)

	values := make([]int64, tuple.Size)
	root.Add(values)
}
