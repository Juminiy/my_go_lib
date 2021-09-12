package algo_compile

import "github.com/Juminiy/my_go_lib/my-algo/algo_compile/fa"

type IFA interface {
	Delta(startState *fa.ISet, edgeA *fa.ISet) *fa.ISet
}
