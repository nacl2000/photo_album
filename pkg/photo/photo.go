package photo

import (
	"os"
	"strconv"
	"path/filepath"
	"os/exec"
	"encoding/json"
	"io/ioutil"
	"github.com/nacl2000/photo_album/pkg/file"
	"github.com/nacl2000/photo_album/pkg/path"
) 

var photoConfigPath = path.GetSourceCodePath("config/photo/photo.json")

type PhotoConfig struct {
	OriginPhotoStoreRoot string `json:"origin_photo_store_root"`  
	CompressPhotoStoreRoot string `json:"compress_photo_store_root"`  
	CompressQuality int `json:"compress_quality"`
}

func readPhotoConfig(photoConfig *PhotoConfig) error {
	data, err := ioutil.ReadFile(photoConfigPath)  
	if err != nil {  
		return err
	}

	if err = json.Unmarshal(data, &photoConfig); err != nil { 
		return err
	}
	return nil
}

func GetPhotoStoreRootPath(location string, is_compress bool) (string, error) {
	var photoConfig PhotoConfig
	if err := readPhotoConfig(&photoConfig); err != nil { 
		return "", err
	}

	if is_compress {
		return filepath.Join(photoConfig.CompressPhotoStoreRoot, location), nil
	}
	return filepath.Join(photoConfig.OriginPhotoStoreRoot, location), nil
}

func GetPhotoStorePath(location string, photoName string, is_compress bool) (string, error) {
	photoStoreRoot, err := GetPhotoStoreRootPath(location, is_compress)
	if err != nil {
		return "", err
	}
	return filepath.Join(photoStoreRoot, photoName), nil
}

func CompressPhoto(originPhotoPath string, compressPhotoPath string) error {
	tmpDir := "/tmp/origin_photo"
	if err := file.RemoveDir(tmpDir); err != nil {  
		return err
	}
	if err := file.CopyFile(originPhotoPath, filepath.Join(tmpDir, filepath.Base(compressPhotoPath))); err != nil {
		return err
	}
	var photoConfig PhotoConfig
	if err := readPhotoConfig(&photoConfig); err != nil { 
		return err
	}
	if err := os.MkdirAll(filepath.Dir(compressPhotoPath), os.ModePerm); err != nil {
		return nil
	}
	compressCmd := exec.Command("collie", "-r", tmpDir, "-o", filepath.Dir(compressPhotoPath), "-q", strconv.Itoa(photoConfig.CompressQuality))
	_, err := compressCmd.Output()
	if err != nil {  
		return err
	}
	return nil
}
