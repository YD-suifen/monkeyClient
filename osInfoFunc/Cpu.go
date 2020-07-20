package osInfoFunc

import (
	"bufio"
	"encoding/json"
	"fmt"
	"monkeyClient/dao"
	"monkeyClient/logUtils"
	"monkeyClient/messageChan"
	"monkeyClient/procFile"
	"os"
	"strconv"
	"strings"
)

var IdleCpuTime float64
var UsedCpuTime float64
var TotalCpuTime float64

type SecondData struct {
	IdleCpuTime float64
	UsedCpuTime float64
}

func GetCpu() osCpu {
	logUtils.Info("GetCpu start")
	var a SecondData
	var b osCpu
	var c dao.SHCpuTable
	a.Update()
	b.Used = Decimal(a.UsedCpuTime)
	b.Idle = Decimal(a.IdleCpuTime)
	c.UsedCpu = b.Used
	c.IdleCpu = b.Idle
	c.HostName = HostName
	c.PrivateIP = PrivateIP
	c.TimeUnix = AtomicClockUnix
	c.KeyName = HostKeyName
	jsonData, _ := json.Marshal(c)
	messageChan.CpuInfo <- jsonData
	return b
}

func (c *SecondData) Update() {
	bI := IdleCpuTime
	bT := TotalCpuTime
	nowTotal, nowIdle := readProc()
	U,I := calculation(nowTotal,nowIdle,bT,bI)
	c.UsedCpuTime = U
	c.IdleCpuTime = I
}

func readProc() (float64,float64) {
	file, err := os.Open(procFile.ReadProc("cpu"))
	if err != nil{
		logUtils.Errorf("GetCpu readProc error=%v",err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	data, _, _ := r.ReadLine()
	dataD := strings.Fields(string(data))

	TotalCpuTime = 0.0
	IdleCpuTime = 0.0

	for _, v := range dataD[1:]{
		iDate,_ := strconv.Atoi(v)
		TotalCpuTime += float64(iDate)
	}
	idles,_ := strconv.Atoi(dataD[4])
	IdleCpuTime = float64(idles)
	return TotalCpuTime,IdleCpuTime
}

func calculation(nT, nI, bT, bI float64) (float64,float64){
	totalTime :=  nT - bT
	totalIdleTime := nI - bI
	noIdeleTime := totalTime - totalIdleTime
	UsedCpu := noIdeleTime * 100 / totalTime
	IdleCpu := float64(100) - UsedCpu
	return UsedCpu,IdleCpu
}


func Decimal(value float64) float64  {

	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

//cpu  565376889 2949 155934444 9083080022 387102381 0 9628079 0 0 0