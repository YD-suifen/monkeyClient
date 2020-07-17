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