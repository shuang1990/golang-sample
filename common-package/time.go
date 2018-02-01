package main

import (
	"time"
	"fmt"
)

func main() {
	now := time.Now()
	fmt.Println("now:", now)
	year, month, day := now.Date()
	fmt.Printf("%d-%d-%d\n", year, month, day)

	year = now.Year()
	month = now.Month()
	day = now.Day()
	fmt.Printf("%d-%d-%d\n", year, month, day)

	hour, min, sec := now.Clock()
	fmt.Printf("%d:%d:%d\n", hour, min, sec)

	hour = now.Hour()
	min = now.Minute()
	sec = now.Second()
	fmt.Printf("%d:%d:%d\n", hour, min, sec)


	d := time.Date(2018, time.January, 2, 19, 30, 50, 100, time.UTC)
	fmt.Printf("%s after %s: %v\n", now.String(), d.String(), now.After(d))


	sts := time.Now().Unix()
	nts := time.Now().UnixNano()
	fmt.Println("current timestamp:", sts, nts)


	current_time := now.Format("2006-01-02 15:04:05")
	fmt.Println("time format: ", current_time)

	t := "2018-01-20 15:20:30"
	date, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("convert %s to timestamp: %d\n", t, date.Unix())

	afterDate := now.Add(time.Hour*24).Format("2006-01-02 15:04:05")
	fmt.Println(afterDate)

	afterDate = now.AddDate(0, 0, 1).Format("2006-01-02 15:04:05")
	fmt.Println(afterDate)
}