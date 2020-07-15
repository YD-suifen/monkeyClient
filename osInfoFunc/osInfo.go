package osInfoFunc

import (
	"monkeyClient/logUtils"
	"monkeyClient/messageChan"
	"net"
	"os"
	"time"
)

var AtomicClockUnix int64
//func init()  {
//	AtomicClockUnix = time.Now().Unix()
//}



func ColleData()  {
	//bbb := logUtils.SugarLogger

	var Host HostInfo
	for  {
		AtomicClockUnix = time.Now().Unix()
		logUtils.Infof("Host.Update time= %v",AtomicClockUnix)
		Host.Update()
		messageChan.HostNowData <- Host
		time.Sleep(time.Second * 60)
	}

}

func (c *HostInfo) Update() {

	c.Times = AtomicClockUnix
	c.HostName = getHostName()
	c.PrivateIP = getIp()
	c.Cpus = GetCpu()
	c.Mems = GetMem()
	c.Disks = GetDisk()
	c.NetConn = GetTcpNet()
}



func getIp() string {
	logUtils.Info("get ip start")
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		logUtils.Errorf("getIp error=%v",err)
		return ""
	}
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				// fmt.Println(ipnet.IP.String())
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getHostName() string {
	logUtils.Info("get hostname start")
	name, err := os.Hostname()
	if err != nil {
		logUtils.Error(err)
		return ""
	}
	return name
}