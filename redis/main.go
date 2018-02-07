package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	redis.DialConnectTimeout(time.Second * 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("set", "name", "zhangshuang")
	if err != nil {
		fmt.Println("set data failed: ", err)
		return
	}

	exists, _ := redis.Bool(conn.Do("exists", "name"))
	fmt.Println(exists)

	_, err = redis.DoWithTimeout(conn, time.Second * 30, "set", "age", 27)
	if err != nil {
		fmt.Println("set data failed: ", err)
		return
	}
}
