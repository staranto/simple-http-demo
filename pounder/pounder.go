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
		url := fmt.Sprintf("http://counter:8181/inc/%s/%s/%d", host, os.Args[1], instanceCount)
		http.Get(url)
		time.Sleep(time.Millisecond * 500)

		log.Printf("host=%s color=%s instanceCount=%d", host, os.Args[1], instanceCount)
	}
}
