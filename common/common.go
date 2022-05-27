package common

const ValuesSize = 6

type ValuesArray [ValuesSize]int64

type Mapper interface {
	Add(total []int64)
}
