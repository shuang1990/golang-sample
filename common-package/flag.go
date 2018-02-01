package main

import (
	"flag"
	"fmt"
)

var (
	config string
)

func main() {
	flag.StringVar(&config, "c", "/etc/nginx/nginx.conf", "please input the nginx config")
	addr := flag.String("addr", "127.0.0.1", "http service address")
	port := flag.Int("port", 8080, "http service listen port")

	flag.Parse()

	fmt.Printf("addr : %s\n", *addr)
	fmt.Printf("config file : %s\n", config)
	fmt.Printf("port: %d\n", *port)
}
