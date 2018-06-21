package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	includeAllFile        bool
	separateByCommas      bool
	showListFormat        bool
	writeSlashIfDirectory bool
)

func initFlag() {
	flag.BoolVar(&includeAllFile, "a", false, "Include directory entries whose names begin with a dot (.).")
	flag.BoolVar(&separateByCommas, "m", false, "Stream output format; list files across the page, separated by commas.")
	flag.BoolVar(&showListFormat, "l", false, "List in long format.")
	flag.BoolVar(&writeSlashIfDirectory, "p", false, "Write a slash (`/') after each filename if that file is a directory.")
}

func ls(dir string) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}

	list := []string{}
	for _, fileInfo := range fileInfos {
		if !includeAllFile && strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}

		filename := fileInfo.Name()
		if writeSlashIfDirectory && fileInfo.IsDir() {
			filename += "/"
		}

		if showListFormat && !separateByCommas {
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

	if separateByCommas {
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
