package osInfoFunc

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"monkeyClient/dao"
	"monkeyClient/logUtils"
	"monkeyClient/messageChan"
	"monkeyClient/procFile"
	"strings"
)

const (
	TCP_ESTABLISHED = "01"
)

type Tcpnet struct {
	TCP_TOTAL int
	TCP_ESTABLISHED int
	TCP_TIME_WAIT int

}

func GetTcpNet() osNetConn {
	logUtils.Info("GetTcpNet start")
	var a osNetConn
	a.Get()
	return a

}


func (c *osNetConn) Get()  {
	contents, err := ioutil.ReadFile(procFile.ReadProc("netconn"))
	if err != nil {
		logUtils.Errorf("osNetConn read file err= %v",err)
		return
	}
	read := bufio.NewReader(bytes.NewBuffer(contents))
	tcpnet := &Tcpnet{}

	for {
		data, _, err := read.ReadLine()
		if err != nil {
			logUtils.Error(err)
			break
		}

		fields := strings.Fields(string(data))
		if fields[0] != "sl"{
			tcpnet.TCP_TOTAL++

			switch string(fields[3][0]) + string(fields[3][1]) {
			case "01":
				tcpnet.TCP_ESTABLISHED++
			case "06":
				tcpnet.TCP_TIME_WAIT++
			}
		}

	}
	c.AllConn = tcpnet.TCP_TOTAL
	c.Established = tcpnet.TCP_ESTABLISHED
	var a dao.SHTcpNetTable
	a.HostName = HostName
	a.PrivateIP = PrivateIP
	a.TimeUnix = AtomicClockUnix
	a.AllConn = c.AllConn
	a.Established = c.Established
	a.KeyName = HostKeyName
	jsonData, _ := json.Marshal(a)
	messageChan.TcpNetInfo <- jsonData
}