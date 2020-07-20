package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"monkeyClient/logUtils"
	"monkeyClient/osInfoFunc"
	"monkeyClient/router"
	"monkeyClient/server"
	"monkeyClient/utils"
	"os"
)

var (
	H bool
	C string
)
func init() {
	flag.BoolVar(&H, "h", false, "help")
	flag.StringVar(&C, "c", "", "read file path!")
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `
Usage: go-execl [-h] [-c config] [-o savedir]

Options:
`)
	flag.PrintDefaults()
}

func main() {

	logUtils.InitLogger()
	logUtils.SugarLogger.Info("log init ok",)


	flag.Parse()
	if H {
		flag.Usage()
		return
	}
	utils.InitRedisConfigs(C)
	logUtils.SugarLogger.Info("config init ok")


	fmt.Println("aaa",utils.Config.Master,utils.Config.Port,utils.Config.KeyName)

	logUtils.SugarLogger.Info("host ColleData start")
	go osInfoFunc.ColleData(utils.Config.KeyName)

	server.SdataDBInsert()

	r := gin.Default()
	router.RegistRouter(r)
	r.Run(":9533")

}