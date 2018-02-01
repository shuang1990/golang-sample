package main

import (
	"os"
	"fmt"
	"strings"
	"io"
)

func main() {
	env := os.Environ()
	for _, val := range env {
		data := strings.Split(val, "=")
		fmt.Printf("%s = %s\n", data[0], data[1])
	}
	hostname, _ := os.Hostname()
	fmt.Println(hostname)

	path := os.Getenv("PATH")
	fmt.Printf("PATH = %s\n", path)

	getFileInfo()
	readFile()
	writeFile()
}

func getFileInfo()  {
	file, err := os.Open("../datasets/example.txt")
	if err != nil {
		fmt.Println("open file failed: ", err)
		return
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info failed: ", err)
		return
	}

	name := fileInfo.Name()
	size := fileInfo.Size()
	fmt.Printf("name:%s, size:%d\n", name, size)


	fileMode := fileInfo.Mode()
	fmt.Printf("is dir: %v\n", fileMode.IsDir())
	fmt.Printf("is file: %v\n", fileMode.IsRegular())
}

func readFile()  {
	file, err := os.Open("../datasets/example.txt")
	if err != nil {
		fmt.Println("open file failed: ", err)
		return
	}

	defer file.Close()

	buf := make([]byte, 1024)

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed: ", err)
			break
		}

		fmt.Println(string(buf[:n]))
	}

	//set file pointer to start then to the offset 100
	_, err = file.Seek(100, 0)
	if err != nil {
		fmt.Println("set seek failed: ", err)
		return
	}

	n, err := file.Read(buf)
	fmt.Println(string(buf[:n]))
}

func writeFile()  {
	file, err := os.OpenFile("../datasets/access.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file with write failed: ", err)
		return
	}
	defer file.Close()

	content := "This is a test message\n"
	_, err = file.Write([]byte(content))
	if err != nil {
		fmt.Println("write file failed: ", err)
		return
	}

	_, err = file.WriteString("Hello world\n")
	if err != nil {
		fmt.Println("write file failed: ", err)
		return
	}
}