package byvalue_test

import (
	"testing"

	"github.com/marcelo-rocha/profiling-pointer-receivers/byvalue"
	"github.com/marcelo-rocha/profiling-pointer-receivers/common"
)

const (
	RootLength  = 100
	InnerLength = 100
)

func BenchmarkByValue(b *testing.B) {
	initialValues := common.ValuesArray{1, 1, 1, 1, 1, 1}
	var root common.Mapper = byvalue.NewInnerNode(initialValues,
		byvalue.NewInnerNode(initialValues,
			byvalue.NewInnerNode(initialValues,
				byvalue.NewInnerNode(initialValues,
					byvalue.NewLeafNode(initialValues),
					InnerLength),
				InnerLength),
			InnerLength),
		RootLength)

	values := make([]int64, common.ValuesSize)
	root.Add(values)
}
