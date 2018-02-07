package main

import (
	"sync"
	"time"
	"math/rand"
	"fmt"
)

var wg sync.WaitGroup
var counter int
var lock sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementor(s string)  {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		lock.Lock()
		counter++
		fmt.Println(s, i, "Counter:", counter)
		lock.Unlock()
	}
	wg.Done()
}
