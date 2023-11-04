package homepage

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/nacl2000/photo_album/pkg/network"
)

func homepageHandler(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("username") == nil {
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
