package files

import "os"

// Проверка существования файла
func Exists(file string) bool {
	fp, err := os.Open(file)
	if err != nil {
		return false
	}
	fp.Close()
	return true
}
