package osInfoFunc

import (
	"github.com/shirou/gopsutil/disk"
	"fmt"
	"monkeyClient/logUtils"
	"strconv"
	"strings"
)

type DevInfo struct{
	DevName string
	PUsed float64
	PTotel uint64
	PFree uint64
}


func decimal(value float64) float64  {

	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func GetDisk() []osDisk {
	logUtils.Info("GetDisk start")

	diskList, _ := disk.Partitions(true)
	var diskInfos osDisk
	var diskInfosList []osDisk

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
		diskInfos.Used = decimal(data.PUsed)
		diskInfos.Free = float64(data.PFree)

		diskInfosList = append(diskInfosList,diskInfos)

		//fmt.Println(data.DevName,data.PTotel,data.PUsed,data.PFree)
		logUtils.Debugf("%v %v %v %v ",data.DevName,data.PTotel,data.PUsed,data.PFree)

	}
	return diskInfosList

}