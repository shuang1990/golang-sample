package main

import (
	"time"
	"fmt"
)

func query(ch chan bool)  {
	time.Sleep(time.Second * 5)
	ch <- true
}

func main() {
	ch := make(chan bool)
	go query(ch)

	t := time.NewTicker(time.Second * 10)

	select{
		case <- t.C:
			fmt.Println("timeout")
		case <-ch:
			fmt.Println("query finished")
	}
}
