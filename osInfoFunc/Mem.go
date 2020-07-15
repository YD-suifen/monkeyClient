package osInfoFunc

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"monkeyClient/logUtils"
	"monkeyClient/procFile"
	"strconv"
	"strings"
)




type Mem struct{
	Buffers uint64
	Cached uint64
	MemTotal uint64
	MemFree uint64
	MemAvailable uint64
}

var multi uint64 = 1024

var Want = map[string]struct{}{
	"Buffers:":      struct{}{},
	"Cached:":       struct{}{},
	"MemTotal:":     struct{}{},
	"MemFree:":      struct{}{},
	"MemAvailable:": struct{}{},
}

func GetMem() osMem {
	logUtils.Info("GetMem start")
	var a osMem
	a.Get()
	return a

}

func (c *osMem) Get()  {

	total,used,free := memInfo()
	pmemFree := float64(free) * 100.0 / float64(total)
	pmemUsed := float64(used) * 100.0 / float64(used)
	c.Used = pmemUsed
	c.Total = float64(total) / (1024 * 2)
	c.Free = pmemFree

}

func (c *Mem) JsonData() {
	data, err := json.Marshal(c)
	if err != nil{
		fmt.Println("ss",err)
		return
	}
	fmt.Println(string(data))
}

func memInfo() (uint64,uint64,uint64) {

	contents, err := ioutil.ReadFile(procFile.ReadProc("mem"))
	if err != nil{
		logUtils.Errorf("memInfo read file err= %v",err)
		//return nil
	}
	read := bufio.NewReader(bytes.NewBuffer(contents))
	mem := &Mem{}

	for  {
		data, _, err := read.ReadLine()
		if err != nil{
			logUtils.Errorf("memInfo error=%v",err)
			break
		}
		fields := strings.Fields(string(data))

		fieldName := fields[0]
		//fmt.Println(fieldName)
		_, ok := Want[fieldName]
		if ok && len(fields) == 3 {

			val, numerr := strconv.ParseUint(fields[1], 10, 64)
			if numerr != nil{
				continue
			}
			switch fieldName {
			case "Buffers:":
				mem.Buffers = val * multi
			case "Cached:":
				mem.Cached = val * multi
			case "MemTotal:":
				mem.MemTotal = val * multi
			case "MemFree:":
				mem.MemFree = val * multi
			case "MemAvailable:":
				mem.MemAvailable = val * multi
			}

		}
	}

	memFree := mem.MemFree + mem.Buffers + mem.Cached

	memUsed := mem.MemTotal - memFree
	return mem.MemTotal,memUsed,memFree

	
}






