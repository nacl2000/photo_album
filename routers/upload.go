package routers

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"github.com/nacl2000/photo_album/pkg/auth"
	"github.com/nacl2000/photo_album/pkg/network"
	"github.com/nacl2000/photo_album/pkg/photo"
)


func upload(c *gin.Context) {
	if err := auth.CheckSeesion(c); err != nil {
		c.Redirect(http.StatusFound, fmt.Sprintf("%s/login", network.GetRequestDomain(c)))
		return
	}
	file, err := c.FormFile("image")
	if err != nil {  
		c.String(http.StatusBadRequest, fmt.Sprintf("获取文件失败: %s", err.Error()))  
		return  
	}
	location := c.PostForm("location")
	originPhotoStoreRootPath, err := photo.GetPhotoStoreRootPath(location, false)
	if err != nil {  
		c.String(http.StatusInternalServerError, "Could not get origin photo store root path.")
		return  
	}
	compressPhotoStoreRootPath, err := photo.GetPhotoStoreRootPath(location, true)
	if err != nil {  
		c.String(http.StatusInternalServerError, "Could not get compress photo store root path.")
		return  
	}
	currentTime := time.Now()
	photoName := fmt.Sprintf("%s.jpg", currentTime.Format("20060102150405"))
	originPhotoPath := filepath.Join(originPhotoStoreRootPath, photoName)
	compressPhotoPath := filepath.Join(compressPhotoStoreRootPath, photoName)
	if err := c.SaveUploadedFile(file, originPhotoPath); err != nil { 
		c.String(http.StatusInternalServerError, fmt.Sprintf("Storing origin photo failed: %s", err.Error()))
	}
	if err := photo.CompressPhoto(originPhotoPath, compressPhotoPath); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Storing compress photo failed: %s", err.Error()))
	}
}

func AddUploadRoutes(router *gin.RouterGroup) {
	loginRouter := router.Group("/upload")
	loginRouter.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.html", "upload")
	})
	loginRouter.POST("/api", upload)
}
