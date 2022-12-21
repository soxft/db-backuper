package tool

import (
	"errors"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		return false
	}
	return false
}

// DeleteLocal delete oldest file if there are more than maxNum
func DeleteLocal(path string, maxNum int) error {
	// 读取 path 下的文件列表
	if fileList, err := os.ReadDir(path); errors.Is(err, nil) {
		for _, file := range fileList {
			if getDirFileNum(path) <= maxNum {
				break
			}
			_ = os.Remove(path + file.Name())
		}
	} else {
		return err
	}
	return nil
}

func getDirFileNum(path string) int {
	if list, err := os.ReadDir(path); errors.Is(err, nil) {
		return len(list)
	}
	return 0
}
