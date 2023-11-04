package main

import (
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/nacl2000/photo_album/routers/photo"
	"github.com/nacl2000/photo_album/routers/login"
	"github.com/nacl2000/photo_album/pkg/path"
)

var router = gin.Default()

func main() {
	getRoutes()
	router.Run(":8080")
}

func rootHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome.html", gin.H{})
}

func getRoutes() {
	htmlPathPattern := filepath.Join("frontend", "*.html")
	router.LoadHTMLGlob(path.GetSourceCodePath(htmlPathPattern))
	cookieStore := sessions.NewCookieStore([]byte("cloud_storage"))
	cookieStore.Options(sessions.Options{
		Path:    "/",
		MaxAge:  1800,
	})
	router.Use(sessions.Sessions("photo_album", cookieStore))

	v1 := router.Group("/")
	v1.GET("/", rootHandler)
	photo.AddRoutes(v1)
	login.AddLoginRoutes(v1)
}
