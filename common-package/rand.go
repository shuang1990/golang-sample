package main

import (
	"math/rand"
	"time"
	"fmt"
)

var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func main() {
	rand.Seed(time.Now().UnixNano())
	//random create 10 integer(0 <= x < 100)
	for i := 0; i < 10; i++ {
		fmt.Printf("round %d: %d\n", i + 1, rand.Intn(100))
	}
	randStr := getRandStr(16)
	fmt.Println(randStr)
}

func getRandStr(num int) string {
	l := len(letter)
	result := make([]rune, num)
	for i := 0; i < num; i++{
		offset := rand.Intn(l)
		result[i] = letter[offset]
	}
	return string(result)
}
