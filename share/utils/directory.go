package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func FindWorkDir() (string, error) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("cannot find work dir")
	}
	for {
		path := filepath.Dir(file)
		base := filepath.Base(path)
		if base == "worframe" {
			return path, nil
		}
		if path == "/" || path == "." {
			// 已经到达了路径的顶部，但是没有找到
			return "", fmt.Errorf("cannot find 'workframe' in path")
		}
		// 继续搜索上一级目录
		file = path
	}

}
