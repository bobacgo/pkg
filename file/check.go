package file

import "os"

// Exists 文件是否存在
func Exists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}