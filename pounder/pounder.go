package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var instanceCount int64

func main() {
	host, _ := os.Hostname()
	for {
		instanceCount++
		node := os.Getenv("NODE")
		if node == "" {
			node = "unknown"
		}
		url := fmt.Sprintf("http://counter:8181/inc/%s/%s/%s/%d", host, node, os.Args[1], instanceCount)
		http.Get(url)
		time.Sleep(time.Millisecond * 250)

		log.Printf("host=%s color=%s instanceCount=%d", host, os.Args[1], instanceCount)
	}
}
