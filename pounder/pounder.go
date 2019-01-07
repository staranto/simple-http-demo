package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var instanceCount int64

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/ready", ready)
	r.HandleFunc("/live", ready)
	go http.ListenAndServe(":8282", r)

	host, _ := os.Hostname()
	for {
		instanceCount++
		node := os.Getenv("NODE")
		if node == "" {
			node = "unknown"
		}
		url := fmt.Sprintf("%s/inc/%s/%s/%s/%d", os.Args[1], host, node, os.Args[2], instanceCount)
		http.Get(url)

		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		sleep := r.Intn(250)
		time.Sleep(time.Millisecond * time.Duration(sleep))

		log.Printf("host=%s color=%s instanceCount=%d sleep=%dms", host, os.Args[1], instanceCount, sleep)
	}
}

func ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
