package osInfoFunc

import (
	"encoding/json"
	"github.com/shirou/gopsutil/disk"
	"monkeyClient/dao"
	"monkeyClient/logUtils"
	"monkeyClient/messageChan"
	"strings"
)

type DevInfo struct{
	DevName string
	PUsed float64
	PTotel uint64
	PFree uint64
}




func GetDisk() []osDisk {
	logUtils.Info("GetDisk start")

	diskList, _ := disk.Partitions(true)
	var diskInfos osDisk
	var diskInfosList []osDisk
	var c dao.SHDiskTable

	data := &DevInfo{}

	for _,v := range diskList{
		if !strings.Contains(v.Device,"/dev"){
			continue
		}

		diskInfo, _  := disk.Usage(v.Mountpoint)

		data.DevName = v.Device
		data.PUsed = diskInfo.UsedPercent
		data.PTotel = diskInfo.Total / (1024 * 1024 * 1024)
		data.PFree = diskInfo.Free / (1024 * 1024 * 1024)
		diskInfos.DevName = data.DevName
		diskInfos.Total = float64(data.PTotel)
		diskInfos.Used = Decimal(data.PUsed)
		diskInfos.Free = float64(data.PFree)

		diskInfosList = append(diskInfosList,diskInfos)
		logUtils.Debugf("%v %v %v %v ",data.DevName,data.PTotel,data.PUsed,data.PFree)

	}
	jsonData, _ := json.Marshal(diskInfosList)
	c.HostName = HostName
	c.PrivateIP = PrivateIP
	c.TimeUnix = AtomicClockUnix
	c.Disk = string(jsonData)
	c.KeyName = HostKeyName
	jsonDataC, _ := json.Marshal(c)

	messageChan.DiskInfo <- jsonDataC
	return diskInfosList

}