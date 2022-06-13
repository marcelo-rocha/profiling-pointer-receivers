package domain

func MakeCollection5(innerFactory InnerMapperFactory, leafFactory LeafMapperFactory) Mapper {
	InitialSlice := make([]int64, ValuesQty)
	for i := range InitialSlice {
		InitialSlice[i] = 1
	}

	return innerFactory(InitialSlice,
		innerFactory(InitialSlice,
			innerFactory(InitialSlice,
				innerFactory(InitialSlice,
					leafFactory(InitialSlice), InnerIterationsQtyL5/8),
				InnerIterationsQtyL5/4),
			InnerIterationsQtyL5/2),
		InnerIterationsQtyL5)
}

func MakeCollection3(innerFactory InnerMapperFactory, leafFactory LeafMapperFactory) Mapper {
	InitialSlice := make([]int64, ValuesQty)
	for i := range InitialSlice {
		InitialSlice[i] = 1
	}

	return innerFactory(InitialSlice,
		innerFactory(InitialSlice,
			leafFactory(InitialSlice), InnerIterationsQtyL3/2),
		InnerIterationsQtyL3)
}

func MakeCollection2(innerFactory InnerMapperFactory, leafFactory LeafMapperFactory) Mapper {
	InitialSlice := make([]int64, ValuesQty)
	for i := range InitialSlice {
		InitialSlice[i] = 1
	}

	return innerFactory(InitialSlice,
		leafFactory(InitialSlice), InnerIterationsQtyL2)

}
