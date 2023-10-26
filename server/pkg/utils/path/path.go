package pathutil

import (
	"os"
	"path/filepath"
)

func GetAbsolutePath(relativePath string) (string, error) {
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func GetRootPath() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			break
		}

		currentDir = filepath.Dir(currentDir)
		if currentDir == filepath.Dir(currentDir) {
			break
		}
	}

	return currentDir, nil
}
