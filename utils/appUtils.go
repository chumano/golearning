package utils

import (
	"fmt"
	"os"
)

func GetWebDirPath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(0)
	}
	fmt.Println(path)
	webDir := path + "/web/dist"

	return webDir
}
