package compile_fa

import "github.com/Juminiy/my_go_lib/my-algo/compile_fa/set_fa"

type IFA interface {
	Delta(startState *set_fa.LimitSet, edgeA *set_fa.InputSet) *set_fa.LimitSet
}
