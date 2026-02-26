package main

import (
	"io/ioutil"
	"fmt"
	"path"
	"strings"
)

func main() {
	ReadDirAll("./")
}

func ReadDirAll(dirName string) {
	fileInfos, _ := ioutil.ReadDir(dirName)
	//fmt.Println(dirName , "----------------------")

	count := strings.Count(dirName , "/")
	for i, v := range fileInfos {
		temp := strings.Repeat("————" , count)
		fmt.Printf("%s %d.%s(%t) \n", temp , i+1, v.Name(), v.IsDir())

		if v.IsDir() {
			ReadDirAll(path.Join(dirName , v.Name()))
		}
	}

}
