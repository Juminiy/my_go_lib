package finite_automata

import "github.com/Juminiy/my_go_lib/my_algo/algo_base/data_struct/complicated"

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
func (iset *ISet) CheckSelf() bool {
	if iset.CharSet != nil && iset.CharSet.Len() > 0 {
		return true
	} else {
		return false
	}
}
