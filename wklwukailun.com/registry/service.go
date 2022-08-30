package registry

// 抽象服务
type Service struct {
	// 服务名
	Name string `json:"name"`
	// 节点列表
	Nodes []*Node `json:"nodes"`
}

// 单个服务节点的抽象
type Node struct {
	Id     string `json:"id"`
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"`
}

// NodeInfo 单个服务节点的信息, Name为服务名称
type NodeInfo struct {
	ServiceName string `json:"name"`
	Node        `json:"node"`
}
