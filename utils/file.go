package utils

import (
	adventofcode "advent-of-code"
	"os"
	"path/filepath"
	"runtime"
)

func GetContentOfFile(path string) string {
	inputPath, _ := filepath.Abs(path)
	data, err := os.ReadFile(inputPath)
	ErrorCheck(err)
	return string(data)
}

func GetData(example bool) string {
	_, file, _, _ := runtime.Caller(1)
	newTask := adventofcode.GetTask(file)

	var input string
	if example {
		input = newTask.GetExample()
	} else {
		input = newTask.GetInput()
	}

	return GetContentOfFile(input)
}
