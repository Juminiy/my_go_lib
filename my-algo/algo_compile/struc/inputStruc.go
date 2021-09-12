package struc

type EdgeInput struct {
	NodeIValue interface{} `nodeIValue:"json"`
	NodeJValue interface{} `nodeJValue:"json"`
	EdgeValue  interface{} `edgeValue:"json"`
}
