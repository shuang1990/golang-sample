package main

import (
	"strings"
	"fmt"
)

func main() {
	str1 := "hello world"
	str2 := "This is a test message"

	result := strings.Compare(str1, str2)
	if result > 0 {
		fmt.Println("str1 > str2")
	} else if(result == 0) {
		fmt.Println("str1 == str2")
	} else {
		fmt.Println("str1 < str2")
	}

	if strings.Contains(str1, "hello") {
		fmt.Println("str1 contains hello")
	}

	nums := strings.Count(str1, "l")
	fmt.Println("str1 count l: ", nums)

	fmt.Println(strings.Fields(str1))
	fmt.Println(strings.HasPrefix(str1, "hel"))

	ss := []string{"aaa", "bbb", "ccc", "ddd"}
	fmt.Println(strings.Join(ss, "-"))

	unique := "123fdsfds-232323-dfdsfsdf-324234"
	res := strings.Split(unique, "-")
	fmt.Println(res)


	reader := strings.NewReader(unique)
	fmt.Printf("%s is %d bytes, size is %d\n", unique, reader.Len(), reader.Size())

	buf := make([]byte, 20)
	n, err := reader.Read(buf)
	if err != nil {
		fmt.Println("read failed: ", err)
		return
	}

	fmt.Println(string(buf[:n]))

}
