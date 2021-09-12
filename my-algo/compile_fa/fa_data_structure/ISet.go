package fa_data_structure

import "github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/complicated"

type ISet struct {
	CharSet *complicated.MySet
}

type DFATable struct {
	State       *ISet
	Input       *ISet
	ChangeTable [][]interface{}
}
