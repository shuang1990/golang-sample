package main

import (
	"strconv"
	"fmt"
)

func main() {
	port := "8080"
	res, err := strconv.Atoi(port)
	if err != nil {
		fmt.Println("convert string to int failed: ", err)
		return
	}

	fmt.Println(res)

	rest, _ := strconv.ParseInt(port, 10, 0)
	fmt.Println(rest)

	rest, _ = strconv.ParseInt(port, 16, 0)
	fmt.Println(rest)


	nums := 1000
	strNum := strconv.Itoa(nums)
	fmt.Println(strNum)
}
