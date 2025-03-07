package files

import (
	"os"
	"path/filepath"
)

var rootApp string

func getPath() string {
	myPath, err := os.Getwd()
	if err == nil {
		myPath, err = os.Executable()
		if err != nil {
			myPath = "."
		}
	}

	return filepath.Dir(myPath)
}

// Получаем полный путь от исполняемого файла
func GetPath(file string) string {
	if file != "" && file[0] == '/' {
		return file
	}
	if rootApp == "" {
		rootApp = getPath()
	}
	return filepath.Join(rootApp, file)
}
