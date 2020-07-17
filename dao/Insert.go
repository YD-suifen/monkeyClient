package dao

import (
	"encoding/json"
	"fmt"
	"monkeyClient/logUtils"
	"monkeyClient/messageChan"
	"monkeyClient/utils"
)

func Insert()  {
	logUtils.Info("Insert start")
	go insertCpu()
	go insertMem()
	go insertDisk()
	go insertTcpNet()
}

func insertCpu() {

	for  {
		select {
		case jsonData := <- messageChan.CpuInfo:

			var data SHCpuTable
			_ = json.Unmarshal(jsonData,&data)
			db := utils.SqlxCli()
			sql := fmt.Sprintf("insert into monkey_s_cpudata (hostName,privateIp,usedCpu,idleCpu,timeUnix) value ('%v','%v',%v,%v,%v)",data.HostName,data.PrivateIP,data.UsedCpu,data.IdleCpu,data.TimeUnix)

			if _, err := db.Exec(sql); err != nil {
				logUtils.Errorf("InsertCpu time=%v,error=%v",data.TimeUnix,err)
			}
			db.Close()
		}
	}

}


func insertMem() {

	for  {
		select {
		case jsonData := <- messageChan.MemInfo:

			var data SHMemTable
			_ = json.Unmarshal(jsonData,&data)
			db := utils.SqlxCli()
			sql := fmt.Sprintf("insert into monkey_s_memdata (hostName,privateIp,total,used,free,timeUnix) value ('%v','%v',%v,%v,%v,%v)",data.HostName,data.PrivateIP,data.Total,data.Used,data.Free,data.TimeUnix)

			if _, err := db.Exec(sql); err != nil {
				logUtils.Errorf("InsertMem time=%v,error=%v",data.TimeUnix,err)
			}
			db.Close()
		}
	}

}

func insertDisk()  {
	for  {
		select {
		case jsonData := <- messageChan.DiskInfo:

			var data SHDiskTable
			_ = json.Unmarshal(jsonData,&data)
			db := utils.SqlxCli()
			sql := fmt.Sprintf("insert into monkey_s_diskdata (hostName,privateIp,disk,timeUnix) value ('%v','%v','%v',%v)",data.HostName,data.PrivateIP,data.Disk,data.TimeUnix)

			if _, err := db.Exec(sql); err != nil {
				logUtils.Errorf("InsertDisk time=%v,error=%v",data.TimeUnix,err)
			}
			db.Close()
		}
	}
}

func insertTcpNet() {

	for  {
		select {
		case jsonData := <- messageChan.TcpNetInfo:

			var data SHTcpNetTable
			_ = json.Unmarshal(jsonData,&data)
			db := utils.SqlxCli()
			sql := fmt.Sprintf("insert into monkey_s_tcpnetdata (hostName,privateIp,allConn,established,timeUnix) value ('%v','%v',%v,%v,%v)",data.HostName,data.PrivateIP,data.AllConn,data.Established,data.TimeUnix)

			if _, err := db.Exec(sql); err != nil {
				logUtils.Errorf("InsertTcpNet time=%v,error=%v",data.TimeUnix,err)
			}
			db.Close()
		}
	}

}