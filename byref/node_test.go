package byref_test

import (
	"testing"

	"github.com/marcelo-rocha/profiling-pointer-receivers/byref"
	"github.com/marcelo-rocha/profiling-pointer-receivers/params"
	"github.com/marcelo-rocha/profiling-pointer-receivers/tuple"
)

func BenchmarkByReference(b *testing.B) {
	initialValues := tuple.TupleOnes
	var root tuple.Mapper = byref.NewInnerNode(initialValues,
		byref.NewInnerNode(initialValues,
			byref.NewInnerNode(initialValues,
				byref.NewInnerNode(initialValues,
					byref.NewLeafNode(initialValues),
					params.InnerIterationsQty),
				params.InnerIterationsQty),
			params.InnerIterationsQty),
		params.RootIterationsQty)

	values := make([]int64, tuple.Size)
	root.Add(values)
}
