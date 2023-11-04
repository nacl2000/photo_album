package network

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetRequestDomain(c *gin.Context) string {
	requestProto := "http"
	if c.Request.Proto == "HTTPS" {  
		requestProto = "https"
	}
	return fmt.Sprintf("%s://%s", requestProto, c.Request.Host)
}
