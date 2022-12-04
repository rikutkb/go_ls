package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	all := flag.Bool("a", false, "隠しファイルの表示を行います。")

	flag.Parse()

	if *all {
		for _, fileInfo := range fileInfos {
			fmt.Println(fileInfo.Name())
		}
	} else {
		for _, fileInfo := range fileInfos {
			if strings.HasPrefix(fileInfo.Name(), ".") {
				continue
			} else {
				fmt.Println(fileInfo.Name())

			}
		}
	}

}
