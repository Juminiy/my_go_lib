package fa

import "github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/complicated"

type ISet struct {
	CharSet *complicated.MySet
}

func (iset *ISet) Construct() {
	mySet := &complicated.MySet{}
	mySet.Construct()
	iset.CharSet = mySet
}

func (iset *ISet) String() string {
	if iset == nil {
		return "nil"
	} else {
		return iset.CharSet.String()
	}
}

type DFATable struct {
	State       *ISet
	Input       *ISet
	ChangeTable [][]interface{}
}
