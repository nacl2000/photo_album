package path

import (
	"os"
	"path/filepath"
)

func GetSourceCodePath(path string) string {
	if os.Getenv("BUILD_WORKSPACE_DIRECTORY") != "" {
		return filepath.Join(os.Getenv("BUILD_WORKSPACE_DIRECTORY"), path)
	}
	return path
}
