package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var instanceCount = 0
var failed = false

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/ready", ready)
	router.HandleFunc("/live", ready)
	go http.ListenAndServe(":8282", router)

	failCeiling := 1000
	if i, err := strconv.Atoi(os.Getenv("FAILCEILING")); err == nil {
		failCeiling = i
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	failAfter := r.Intn(failCeiling)

	node := os.Getenv("NODE")
	if node == "" {
		node = "unknown"
	}
	host, _ := os.Hostname()

	for {
		instanceCount++
		if instanceCount > failAfter {
			failed = true
		}

		url := fmt.Sprintf("%s/inc/%s/%s/%s/%d", os.Args[1], host, node, os.Args[2], instanceCount)
		http.Get(url)

		s = rand.NewSource(time.Now().UnixNano())
		r = rand.New(s)
		sleep := r.Intn(250)
		time.Sleep(time.Millisecond * time.Duration(sleep))

		log.Printf("host=%s color=%s instanceCount=%d failAfter=%d sleep=%dms",
			host, os.Args[1], instanceCount, failAfter, sleep)
	}
}

func ready(w http.ResponseWriter, r *http.Request) {
	if !failed {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
