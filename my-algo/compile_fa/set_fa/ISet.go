package set_fa

type LimitSet struct {
	CharSet []byte
}

type InputSet struct {
	CharSet []byte
}

type DFATable struct {
	State       *LimitSet
	Input       *InputSet
	ChangeTable [][]interface{}
}
