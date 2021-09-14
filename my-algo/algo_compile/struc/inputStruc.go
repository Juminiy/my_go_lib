package struc

type EdgeInput struct {
	NodeIValue interface{} `NodeIValue:"json"`
	NodeJValue interface{} `NodeJValue:"json"`
	EdgeValue  interface{} `EdgeValue:"json"`
}
type ValuesInput struct {
	Edges []EdgeInput   `Edges:"json"`
	Nodes []interface{} `Nodes:"json"`
}
