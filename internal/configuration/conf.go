package configuration

import (
	"fmt"
	"os"
	"runtime"
)

func projectRootDirectory() string {
	var directoryPath string = ""

	switch runtime.GOOS {
	case "windows":
		directoryPath = "C:\\Program Files (x86)\\go-sync-ex\\"
	case "darwin":
		directoryPath = "/opt/go-sync-ex.d"
	case "linux":
		directoryPath = "/opt/go-sync-ex.d"
	default:
		directoryPath = "/opt/go-sync-ex.d"
	}
	return directoryPath
}

func ConfInit() {
	directoryPath := projectRootDirectory()
	var filePath string = directoryPath + "/conf.init"

	err := os.MkdirAll(directoryPath, 0755)

	if err != nil {
		panic(err)
	}

	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Printf("Configuration file %s does not exist\n", filePath)
		fmt.Printf("Create configuration file at %s\n", filePath)
		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		errc := file.Close()
		if errc != nil {
			return
		}
	}
	InitConfig()
}

func InitConfig() {
	filePath := "/opt/go-sync-ex.d/conf.init"

	// Open file and create if not exist
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0744)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Save to file
	_, err = file.WriteString("OS: " + runtime.GOOS + "\n")
	if err != nil {
		panic(err)
	}
}
