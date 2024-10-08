package middleware

import (
	"github.com/gin-gonic/gin"
	"iplist/service"
	"log"
)

func Authentication(c *gin.Context) {
	ip := c.ClientIP()
	server, err := service.ServerSvcIns.GetServerByIp(ip)
	if err != nil {
		log.Println(err)
		c.Next()
		return
	}
	c.Set("server", server)
	c.Next()
}
