package utils

import (
	"path"
	"runtime"
)

// GetFilePath 获取根目录路径
/* @description 获取根目录路径,方便读取根目录下的其他路径
 * @since 2023/2/1915:21
 * @param args[0] 文件名
 * @param args[1] 相对于根目录下的目录路径, 例如: /config/
 * @return 根据传入的参数组成的基于项目路径的绝对路径
 *  */
func GetFilePath(args ...string) string {
	var dir string
	if len(args) < 2 || args[1] == "" {
		dir = "/config/"
	}
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(path.Dir(filename))    //获取当前工作目录
	dirPath := path.Dir(root + dir)         // 获取配置文件的目录
	filePath := path.Join(dirPath, args[0]) // 获取配置文件
	return filePath
}
