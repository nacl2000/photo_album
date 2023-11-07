package routers

import (
	"os"
	"io"
	"encoding/json" 
	"fmt"
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/nacl2000/photo_album/pkg/network"
	"github.com/nacl2000/photo_album/pkg/photo"
)

type Image struct {  
	Url string `json:"url"`  
}

func displayPhotoHandler(c *gin.Context) {
	is_compress := false
	location := c.Query("location")
	photoName := c.Query("photo_name")
	if c.Query("is_compress") != "" {
		is_compress = true
	}
	if photoName == "" {
		c.HTML(http.StatusOK, "display_photo.html", gin.H{"location": location})
		return
	}
	photoPath, err:= photo.GetPhotoStorePath(location, photoName, is_compress)
	if err != nil {  
		c.String(http.StatusInternalServerError, "Could not get photo store path.")
		return  
	}
	photo, err := os.Open(photoPath)
	if err != nil {  
		c.String(http.StatusInternalServerError, fmt.Sprintf("Could not open photo file %s", photoPath))
		return  
	}
	defer photo.Close()  
  
	_, err = io.Copy(c.Writer, photo)  
	if err != nil {  
		c.String(http.StatusInternalServerError, "Could not send photo.") 
		return
	}
	c.Header("Content-Type", "image/jpeg")
}

func getCompressPhotoUrlHandler(c *gin.Context) {
	imageUrls := []Image{ }
	location := c.Query("location")
	phtotRootPath, err := photo.GetPhotoStoreRootPath(location, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	err = filepath.Walk(phtotRootPath, func (path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			imageUrls = append(imageUrls, Image{Url: fmt.Sprintf("%s/photo/display?is_compress=1&photo_name=%s&location=%s", network.GetRequestDomain(c), filepath.Base(path), location)})
			return nil
		}
		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) 
		return
	}
	jsonResponse, err := json.Marshal(imageUrls)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}) 
		return
	}
	c.JSON(http.StatusOK, gin.H{"imageUrls": string(jsonResponse)})
}

func AddRoutes(router *gin.RouterGroup) {
	photoRouter := router.Group("/photo")
	photoRouter.GET("/get_compress_photo_url", getCompressPhotoUrlHandler)
	photoRouter.GET("/display", displayPhotoHandler)
}
