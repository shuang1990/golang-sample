package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main() {
	file, err := os.Open("../datasets/example.txt")
	if err != nil {
		fmt.Println("open file failed: ", err)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("read file failed: ", err)
		return
	}

	fmt.Println(string(data))

	dir := "."

	filesInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("read dir failed: ", err)
		return
	}

	for _, fileInfo := range filesInfo {
		fmt.Printf("%s is %d bytes and mode is %s\n", fileInfo.Name(), fileInfo.Size(), fileInfo.Mode().String())
	}


	data, err = ioutil.ReadFile("bufio.go")
	if err != nil {
		fmt.Println("read file failed: ", err)
		return
	}

	fmt.Println(string(data))

	err = ioutil.WriteFile("../datasets/bufio.bak.go", data, 0644)
	if err != nil {
		fmt.Println("write file failed: ", err)
	}


}
