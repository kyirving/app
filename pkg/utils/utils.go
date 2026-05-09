package utils

import "os"

// IsDirExists 目录是否存在
func IsDirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	// 存在且是目录
	return info.IsDir()
}

// IsFileExists 文件是否存在
func IsFileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	// 存在且是文件
	return info.Mode().IsRegular()
}
