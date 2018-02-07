package main

import (
	"time"
	"fmt"
)

func query(ch chan int)  {
	time.Sleep(time.Second * 5)
}

func main() {
	ch := make(chan int)
	go query(ch)

	t := time.NewTicker(time.Second * 10)

	select{
		case <- t.C:
			fmt.Println("timeout")
		case <-ch:
			fmt.Println("query finished")
	}
}
