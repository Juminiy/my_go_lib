package compile_fa

import "github.com/Juminiy/my_go_lib/my-algo/compile_fa/fa_data_structure"

type IFA interface {
	Delta(startState *fa_data_structure.ISet, edgeA *fa_data_structure.ISet) *fa_data_structure.ISet
}
