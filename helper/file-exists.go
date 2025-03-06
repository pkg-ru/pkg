package helper

import "os"

// Проверка существования файла
func FileExists(file string) bool {
	fp, err := os.Open(file)
	if err != nil {
		return false
	}
	fp.Close()
	return true
}
