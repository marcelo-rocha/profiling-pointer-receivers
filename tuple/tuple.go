package tuple

const Size = 5

type Tuple [Size]int64

type Mapper interface {
	Add(total []int64)
}

var TupleOnes = Tuple{1, 1, 1, 1, 1}
