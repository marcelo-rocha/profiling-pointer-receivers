package tuple

const Size = 5

type Tuple struct {
	V0, V1, V2, V3, V4 int64
}

type Mapper interface {
	Add(total []int64)
}

var TupleOnes = Tuple{1, 1, 1, 1, 1}
