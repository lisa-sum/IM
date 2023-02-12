package utils

import (
	"path"
	"runtime"
)

func GetPath(file, dir string) string {
	if dir == "" {
		dir = "/config/"
	}
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename)) //获取当前工作目录
	dirPath := path.Dir(root + dir)      // 获取配置文件的目录
	filePath := path.Join(dirPath, file) // 获取配置文
	return filePath
}
