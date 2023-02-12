package test

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

func TestGetPath(t *testing.T) {
	// 获取当前文件的路径
	_, filename, _, _ := runtime.Caller(0)
	rootDir := path.Dir(path.Dir(filename))
	fmt.Println("rootDir:", rootDir)
	filePath := GetPath("db.yaml", "")
	fmt.Println("filePath:", filePath)
	fmt.Println("Test Successfully!")
}
