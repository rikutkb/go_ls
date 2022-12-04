package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
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
	var (
		allFlag        bool
		allmostAllFlag bool
		sortTime       bool
		sortReverse    bool
		fullTime       bool
	)
	flag.BoolVar(&allFlag, "a", false, "隠しファイルの表示を行います。")
	flag.BoolVar(&allFlag, "all", false, "隠しファイルの表示を行います。")
	flag.BoolVar(&allmostAllFlag, "A", false, ".,..を除いた隠しファイルの表示を行います。")
	flag.BoolVar(&allmostAllFlag, "allmost-all", false, ".,..を除いた隠しファイルの表示を行います。")
	flag.BoolVar(&sortTime, "t", false, "新しい順に時刻でのソートを行います。")
	flag.BoolVar(&sortReverse, "t", false, "並び順を逆にします。。")
	flag.BoolVar(&fullTime, "full-time", false, "作成時刻の表示を行います。")

	flag.Parse()

	if !allFlag {
		for i, fileInfo := range fileInfos {
			if strings.HasPrefix(fileInfo.Name(), ".") {
				fileInfos = fileInfos[:i+copy(fileInfos[i:], fileInfos[i+1:])]

			}
		}
	}

	// sort
	if sortTime {
		sort.SliceStable(fileInfos, func(i, j int) bool {
			return fileInfos[i].ModTime().After(fileInfos[j].ModTime())
		})
	}
	// if sortReverse {
	// }
	for _, fileInfo := range fileInfos {
		fmt.Print(fileInfo.Name())
		if fullTime {
			fmt.Print(" " + fileInfo.ModTime().String())
		}
		fmt.Println()
	}

}
