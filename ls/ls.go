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
	p bool
)

func ls(dir string) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	list := []string{}
	for _, fileInfo := range fileInfos {
		file := fileInfo.Name()
		if p && fileInfo.IsDir() {
			file += "/"
		}

		list = append(list, file)
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
	flag.BoolVar(&m, "m", false, "Stream output format; list files across the page, separated by commas.")
	flag.BoolVar(&p, "p", false, "Write a slash (`/') after each filename if that file is a directory.")
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
