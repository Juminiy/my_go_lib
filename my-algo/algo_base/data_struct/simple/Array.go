package simple

// Array TODO
type Array struct {
	Slice  []int
	SI, EI int
}

func (array *Array) Construct() {
	array.Slice = make([]int, 0)
	array.SI = 0
	array.EI = 0
}
