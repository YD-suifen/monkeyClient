package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"monkeyClient/messageChan"
)


func Jiange(c *gin.Context)  {
	//name := c.Params("name")
	fmt.Println("hello jiange")

}

func HostInfoGet(c *gin.Context)  {

	for  {
		select {
		case data := <-messageChan.HostNowData:
			jsondata,_ := json.Marshal(data)
			//data :=
			//logUtils.Debug(string(jsondata))
			fmt.Println(string(jsondata))
		}

	}
	
}