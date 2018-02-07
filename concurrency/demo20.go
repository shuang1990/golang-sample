package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.NewTicker(time.Second)
	for v := range t.C{
		fmt.Println(time.Now().Unix())
		fmt.Println(v.Format("2006-01-02 15:04:05"))
	}
}
