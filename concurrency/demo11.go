package main

import (
	"sync"
	"runtime"
	"fmt"
	"time"
)

var wg sync.WaitGroup

func init()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	foo()
	bar()
	wg.Wait()
}

func foo()  {
	for i := 0;i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar()  {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
