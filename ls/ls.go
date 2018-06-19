package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	m bool
)

func ls(dir string) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	list := []string{}
	for _, fileInfo := range fileInfos {
		list = append(list, fileInfo.Name())
	}

	if m {
		fmt.Println(strings.Join(list, ", "))
	} else {
		for _, file := range list {
			fmt.Println(file)
		}
	}
}

func initFlag() {
	flag.BoolVar(&m, "m", false, "separated by commas")
}

func getDir() string {
	if flag.NArg() > 0 {
		return flag.Arg(0)
	} else {
		dir, _ := os.Getwd()
		return dir
	}
}

func main() {
	initFlag()
	flag.Parse()

	dir := getDir()
	ls(dir)
}
