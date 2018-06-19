package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ls(dir string) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}
}

func main() {
	dir, _ := os.Getwd()
	ls(dir)
}
