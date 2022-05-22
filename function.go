package main

import (
	"fmt"
	"io/ioutil"
)

func getFiles(path string) ([]string, error) {
	var files []string
	dirFiles, err := ioutil.ReadDir(fmt.Sprintf("%s", path))
	if err != nil {
		return nil, err
	}
	for _, i := range dirFiles {
		files = append(files, i.Name())
	}
	return files, nil
}


func returnRog() (rog int) {
	rog = 123
	return
}


func main(){
	privet := func(x, y int) int {
		i := (x + y) * 2
		return i
	}



	f, _ := getFiles (".")
	fmt.Println(f)
	fmt.Println(returnRog())
	fmt.Println(privet(1, 5))
}