package messageChan

var (
	HostNowData chan interface{}
)

func init()  {

	HostNowData = make(chan interface{}, 10000)
}