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
	var dir string
	if len(os.Args) > 1 {
		dir = os.Args[1]
	} else {
		dir, _ = os.Getwd()
	}

	ls(dir)
}
