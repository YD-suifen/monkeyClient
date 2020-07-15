package osInfoFunc

type HostInfo struct {
	HostName string `json:"hostName"`
	PrivateIP string `json:"privateIP"`
	Cpus osCpu `json:"cpus"`
	Mems osMem `json:"mems"`
	Disks []osDisk `json:"disks"`
	NetConn osNetConn `json:"netConn"`
	ProccInfo osProcc `json:"proccInfo"`
	Times int64 `json:"times"`
}

type osCpu struct {
	Used float64 `json:"used"`
	Idle float64 `json:"idle"`
}

type osMem struct {
	Total float64 `json:"total"`
	Used float64 `json:"used"`
	Free float64 `json:"free"`
}

type osDisk struct {
	DevName string `json:"devName"`
	Total float64 `json:"total"`
	Used float64 `json:"used"`
	Free float64 `json:"free"`
}

type osNetConn struct {
	AllConn int `json:"allConn"`
	Established int `json:"established"`
}

type osProcc struct {
	ProccName string `json:"proccName"`
	ProccCpu float64 `json:"proccCpu"`
	ProccMem float64 `json:"proccMem"`
}