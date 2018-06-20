package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	a bool
	m bool
	l bool
	p bool
)

func initFlag() {
	flag.BoolVar(&a, "a", false, "Include directory entries whose names begin with a dot (.).")
	flag.BoolVar(&m, "m", false, "Stream output format; list files across the page, separated by commas.")
	flag.BoolVar(&l, "l", false, "List in long format.")
	flag.BoolVar(&p, "p", false, "Write a slash (`/') after each filename if that file is a directory.")
}

func ls(dir string) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	list := []string{}
	for _, fileInfo := range fileInfos {
		if !a && strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}

		filename := fileInfo.Name()
		if p && fileInfo.IsDir() {
			filename += "/"
		}

		if l && !m {
			output := fmt.Sprintf(
				"%v\t%10d\t%v\t%v",
				fileInfo.Mode(),
				fileInfo.Size(),
				fileInfo.ModTime().Format("Jan _2 15:04"),
				filename,
			)
			list = append(list, output)
		} else {
			list = append(list, filename)
		}
	}

	if m {
		fmt.Println(strings.Join(list, ", "))
	} else {
		for _, file := range list {
			fmt.Println(file)
		}
	}
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
