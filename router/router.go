package router

import (
	"github.com/gin-gonic/gin"
	"monkeyClient/server"
)


func RegistRouter(r *gin.Engine)  {

	r.GET("/jiange",server.Jiange)
	r.GET("/hostagent",server.HostInfoGet)

}