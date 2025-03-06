package helper

import "os"

func FileExists(file string) bool {
	fp, err := os.Open(file)
	if err != nil {
		return false
	}
	fp.Close()
	return true
}
