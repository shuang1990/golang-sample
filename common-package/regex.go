package main

import (
	"regexp"
	"fmt"
	"time"
)

func main() {
	pattern := "\\d{4}-\\d{2}-\\d{2}"
	regex, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		return
	}

	today := time.Now().Format("2006-01-02 15:04:05")
	result := regex.FindString(today)
	flag := regex.MatchString(today)
	fmt.Println("match " + today + ": ", result, flag)

	repl := regex.ReplaceAllString(today, "")
	fmt.Println(repl)

	date := "dfssd-fdsfs-fdsd"
	result = regex.FindString(date)
	flag = regex.MatchString(date)
	fmt.Println("match " + date + ": ", result, flag)

	flag, err = regexp.MatchString("\\d{4}-\\d{2}-\\d{2}", "2018-02-03 10:45:40")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("match: ", flag)

	pattern = "-"
	regex, err = regexp.Compile(pattern)
	if err != nil {
		fmt.Println(err)
		return
	}

	res := regex.Split("2018-02-04", -1)
	fmt.Println(res)
}
