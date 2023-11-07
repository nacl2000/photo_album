package file

import (  
	"io"  
	"os"
	"path/filepath"
)


func CopyFile(sourcePath string, destinationPath string) error {
	_, err := os.Stat(sourcePath)
   	if err != nil {
      return err
   	}
	src, err := os.Open(sourcePath)  
	if err != nil {  
		return err  
	}  
	defer src.Close()  
	
	if err := os.MkdirAll(filepath.Dir(destinationPath), os.ModePerm); err != nil {
		return nil
	}

	dest, err := os.Create(destinationPath)
	if err != nil {  
		return err  
	}

	defer dest.Close()  
	  
	_, err = io.Copy(dest, src)  
	if err != nil {  
		return err  
	}  
	 
	sourceInfo, err := os.Stat(sourcePath)  
	if err != nil {  
		return err  
	}  
	err = os.Chmod(destinationPath, sourceInfo.Mode())  
	if err != nil {  
		return err  
	}  
	 
	return nil  
}

func RemoveDir(dirPath string) error {
	if _, err := os.Stat(dirPath); err == nil {  
		if err := os.RemoveAll(dirPath); err != nil {  
			return err
		}
	}
	return nil
}
