package graph_fa

import (
	ds "github.com/Juminiy/my_go_lib/my-algo/algo_base/data_struct/complicated"
	fa "github.com/Juminiy/my_go_lib/my-algo/compile_fa"
)

const (
	epsilon = "epsilon"
)

func RegexToGraph(regex string) *ds.AdjGraph {
	if len(regex) == 0 {
		return nil
	}
	adj := &ds.AdjGraph{}
	adj.Construct(true)
	startNode, q0Node := &ds.GraphNode{"start"}, &ds.GraphNode{"q0"}
	adj.AddNode(startNode)
	adj.AddNode(q0Node)
	// subStr := ""
	for _, ch := range regex {
		switch ch {

		case '(':
			{
				// subStr = " "
				break
			}
		case ')':
			{
				break
			}
		case '|':
			{
				break
			}

		default:
			{

			}
		}
	}
	return adj
}

func GraphToNFA(graph *ds.AdjGraph) *fa.NFA {
	return nil
}

func NFAToDFA(nfa *fa.NFA) *fa.DFA {
	return nil
}
