package test

import (
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestGetPath(t *testing.T) {

	file, err := exec.LookPath(os.Args[0])
	path, err := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("filePath:", file)
	t.Log("path:", path)
	t.Log("index:", index)

}
func TestGetPath1(t *testing.T) {
	WorkPath, err := os.Getwd()
	sysType := runtime.GOOS

	if sysType == "windows" {
		WorkPath = filepath.ToSlash(WorkPath)
	} else if sysType == "linux" {
		WorkPath = filepath.FromSlash(WorkPath)
	}

	if err != nil {
		panic(err)
	}
	t.Log("WorkPath:", path.Clean("../"+WorkPath))
	t.Log("WorkPath1:", path.Base("../"+WorkPath))

	str1 := filepath.FromSlash(WorkPath)
	str2 := filepath.ToSlash(WorkPath)
	log.Println("str1:", str1)
	log.Println("str2:", str2)
}
