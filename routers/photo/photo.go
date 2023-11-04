package photo

import (
	"os"
	"io"
	"encoding/json" 
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"github.com/gin-gonic/gin"
	"github.com/nacl2000/photo_album/models/image_model"
	"github.com/nacl2000/photo_album/pkg/network"
)

var originPhtotRootPath = filepath.Join("/home/nacl", "photo_data/")
var compressPhtotRootPath = filepath.Join("/home/nacl", "compress_photo_data/")

func displayPhotoHandler(c *gin.Context) {
	log.SetOutput(os.Stdout)
	phtotRootPath := originPhtotRootPath
	location := c.Query("location")
	photoName := c.Query("photo_name")
	if c.Query("is_compress") != "" {
		phtotRootPath = compressPhtotRootPath
	}
	log.Println("phtotRootPath:", phtotRootPath)
	log.Println("location:", location)
	log.Println("photoName:", photoName)
	if photoName == "" {
		c.HTML(http.StatusOK, "display_photo.html", gin.H{"location": location})
		return
	}
	photoPath := filepath.Join(phtotRootPath, location, photoName)
	photo, err := os.Open(photoPath)
	if err != nil {  
		c.String(http.StatusInternalServerError, "Could not open photo file")
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
	imageUrls := []image_model.Image{ }
	location := c.Query("location")
	phtotRootPath := filepath.Join(compressPhtotRootPath, location)
	err := filepath.Walk(phtotRootPath, func (path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			imageUrls = append(imageUrls, image_model.Image{Url: fmt.Sprintf("%s/photo/display?is_compress=1&photo_name=%s&location=%s", network.GetRequestDomain(c), filepath.Base(path), location)})
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
