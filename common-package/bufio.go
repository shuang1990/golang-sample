package main

import (
	"bufio"
	"os"
	"fmt"
	"io"
)

func main() {
	//readFromStdin()
	//bufReadFile()
	//bufWriteFile()
	bufCopy()
}

func bufCopy()  {
	input, err := os.Open("../datasets/example.txt")
	if err != nil {
		fmt.Println("open file failed: ", err)
		return
	}

	defer input.Close()

	output, err := os.OpenFile("../datasets/bak.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed: ", err)
		return
	}
	defer output.Close()

	n, err := io.Copy(output, input)
	if err != nil {
		fmt.Println("copy file failed: ", err)
		return
	}

	fmt.Printf("totol copy %d bytes\n", n)

}

func readFromStdin()  {
	fmt.Println("pelase input a string:")
	reader := bufio.NewReader(os.Stdin)
	content, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read failed: ", err)
		return
	}
	fmt.Println(content)
}

func bufWriteFile()  {
	file, err := os.OpenFile("../datasets/buffile.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open file failed: ", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	content := "hello world\n"
	_, err = writer.WriteString(content)

	if err != nil {
		fmt.Println("write file failed: ", err)
		return
	}

	content = "My name is wilson\n"
	_, err = writer.Write([]byte(content))
	if err != nil {
		fmt.Println("write file failed: ", err)
		return
	}

	writer.Flush()
}

func bufReadFile()  {
	file, err := os.Open("../datasets/example.txt")
	if err != nil {
		fmt.Println("open file failed: ", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		_, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed: ", err)
			break
		}
		fmt.Println("sfdsdfs")
	}

	//set file pointer to start
	file.Seek(0, 0)

	//buf := make([]byte, 1024)
	//n, err := reader.Read(buf)
	//fmt.Println(string(buf[:n]))

	for {
		data, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed: ", err)
			break
		}

		fmt.Print(string(data))
	}
}