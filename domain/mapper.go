package domain

type Mapper interface {
	Add(total []int64)
}

type InnerMapperFactory func(values []int64, child Mapper, iterationsQty int) Mapper

type LeafMapperFactory func(values []int64) Mapper

const ValuesQty = 5
