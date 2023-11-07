package routers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/nacl2000/photo_album/pkg/auth"
	"github.com/nacl2000/photo_album/pkg/network"
)

func homepageHandler(c *gin.Context) {
	if err := auth.CheckSeesion(c); err != nil {
		c.Redirect(http.StatusFound, fmt.Sprintf("%s/login", network.GetRequestDomain(c)))
		return
	}
    c.HTML(http.StatusOK, "homepage.html", "login")
	return
}

func AddHomepageRoutes(router *gin.RouterGroup) {
	loginRouter := router.Group("/homepage")
	loginRouter.GET("/", homepageHandler)
}
