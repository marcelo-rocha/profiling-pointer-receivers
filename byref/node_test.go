package byref_test

import (
	"testing"

	"github.com/marcelo-rocha/profiling-pointer-receivers/byref"
	"github.com/marcelo-rocha/profiling-pointer-receivers/common"
)

const (
	RootLength  = 100
	InnerLength = 100
)

func BenchmarkByReference(b *testing.B) {
	initialValues := common.ValuesArray{1, 1, 1, 1, 1, 1}
	var root common.Mapper = byref.NewInnerNode(initialValues,
		byref.NewInnerNode(initialValues,
			byref.NewInnerNode(initialValues,
				byref.NewInnerNode(initialValues,
					byref.NewLeafNode(initialValues),
					InnerLength),
				InnerLength),
			InnerLength),
		RootLength)

	values := make([]int64, common.ValuesSize)
	root.Add(values)
}
