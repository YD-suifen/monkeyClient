package messageChan

var (
	HostNowData chan interface{}
	CpuInfo chan []byte
	MemInfo chan []byte
	DiskInfo chan []byte
	TcpNetInfo chan []byte
)

func init()  {

	HostNowData = make(chan interface{}, 10)
	CpuInfo = make(chan []byte, 10)
	MemInfo = make(chan []byte, 10)
	DiskInfo = make(chan []byte, 10)
	TcpNetInfo = make(chan []byte, 10)
}

func GetChanLen(chanName string) int {
	switch chanName {
	case "HostNowData":
		return len(HostNowData)
	case "CpuInfo":
		return len(CpuInfo)
	case "MemInfo":
		return len(MemInfo)
	case "DiskInfo":
		return len(DiskInfo)
	case "TcpNetInfo":
		return len(TcpNetInfo)
	}
	return 0
}