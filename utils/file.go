package utils

import (
	"os"
	"path/filepath"
)

func GetContentOfFile(path string) string {
	inputPath, _ := filepath.Abs(path)
	data, error := os.ReadFile(inputPath)
	ErrorCheck(error)
	return string(data)
}
