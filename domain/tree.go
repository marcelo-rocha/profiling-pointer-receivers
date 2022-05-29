package domain

func MakeTree(innerFactory InnerMapperFactory, leafFactory LeafMapperFactory) Mapper {
	InitialSlice := make([]int64, ValuesQty)
	for i := range InitialSlice {
		InitialSlice[i] = 1
	}

	return innerFactory(InitialSlice,
		innerFactory(InitialSlice,
			innerFactory(InitialSlice,
				innerFactory(InitialSlice,
					leafFactory(InitialSlice), InnerIterationsQty/8),
				InnerIterationsQty/4),
			InnerIterationsQty/2),
		InnerIterationsQty)
}
