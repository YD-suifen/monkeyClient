package osInfoFunc

import (
	"bufio"
	"fmt"
	"monkeyClient/logUtils"
	"monkeyClient/procFile"
	"os"
	"strconv"
	"strings"
)

var IdleCpuTime float64
var UsedCpuTime float64

type SecondData struct {
	TotalCpuTime float64
	IdleCpuTime float64
}
func GetCpu() osCpu {
	logUtils.Info("GetCpu start")
	var a SecondData
	var b osCpu
	a.Update()
	b.Used = Decimal(UsedCpuTime)
	b.Idle = Decimal(IdleCpuTime)
	return b
}
func readProc() *SecondData {
	file, _ := os.Open(procFile.ReadProc("cpu"))
	defer file.Close()

	r := bufio.NewReader(file)
	data, _, _ := r.ReadLine()
	dataD := strings.Fields(string(data))
	sorData := &SecondData{}

	for _, v := range dataD[1:]{
		iDate,_ := strconv.Atoi(v)
		sorData.TotalCpuTime += float64(iDate)
	}
	idles,_ := strconv.Atoi(dataD[4])
	sorData.IdleCpuTime = float64(idles)
	return sorData
}
func (c *SecondData) Update() {
	if c.TotalCpuTime == 0{
		*c = SecondData{}
	}
	var beData SecondData
	beData.TotalCpuTime = c.TotalCpuTime
	beData.IdleCpuTime = c.IdleCpuTime
	nowData := readProc()
	calculation(*nowData,beData)
}
func calculation(sData,eData SecondData) {
	totalTime := eData.TotalCpuTime - sData.TotalCpuTime
	totalIdleTime := eData.IdleCpuTime - sData.IdleCpuTime
	noIdeleTime := totalTime - totalIdleTime
	UsedCpuTime = noIdeleTime * 100 / totalTime
	IdleCpuTime = float64(100) - UsedCpuTime
}


func Decimal(value float64) float64  {

	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

//cpu  565376889 2949 155934444 9083080022 387102381 0 9628079 0 0 0