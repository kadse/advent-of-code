package adventofcode

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type ITask interface {
	GetExample() string
	GetInput() string
}

type Task struct {
	day  int
	year int
}

func (t Task) GetExample() string {
	return filepath.Join(getPathTask(t), "example")
}

func (t Task) GetInput() string {
	return filepath.Join(getPathTask(t), "input")
}

func GetProjectDirPath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Dir(file)
}

func GetTask(filePath string) ITask {
	relativeTaskPath, _ := filepath.Rel(GetProjectDirPath(), filePath)
	pathParts := strings.Split(relativeTaskPath, string(filepath.Separator))

	day, _ := strconv.Atoi(strings.Replace(pathParts[1], "day", "", 1))
	year, _ := strconv.Atoi(pathParts[0])

	return ITask(*newTask(day, year))
}

func newTask(day int, year int) *Task {
	return &Task{day, year}
}

func getPathTask(t Task) string {
	return filepath.Join(GetProjectDirPath(), strconv.Itoa(t.year), "day"+strconv.Itoa(t.day))
}
