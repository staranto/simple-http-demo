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
var useProbe = false

func main() {

	_, useProbe = os.LookupEnv("PROBE")

	router := mux.NewRouter()
	router.HandleFunc("/ready", ready)
	router.HandleFunc("/live", ready)
	go http.ListenAndServe(":8282", router)

	failFloor := 500
	failCeiling := 1000
	if i, err := strconv.Atoi(os.Getenv("FAILCEILING")); err == nil {
		if i > failCeiling {
			failCeiling = i
		}
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	failAfterCount := r.Intn(failCeiling-failFloor) + failFloor

	node := os.Getenv("NODE")
	if node == "" {
		node = "unknown"
	}
	host, _ := os.Hostname()

	color := os.Getenv("COLOR")
	if color == "" {
		if len(os.Args) > 2 {
			color = os.Args[2]
		}
		if color == "" {
			color = "blue"
		}
	}

	for {
		instanceCount++
		if instanceCount > failAfterCount {
			log.Printf("Fail celing hit.  Using probe signal=%t", useProbe)
			if useProbe {
				failed = true
			} else {
				break
			}
		}

		url := fmt.Sprintf("%s/inc/%s/%s/%s/%d", os.Args[1], host, node, color, instanceCount)
		http.Get(url)

		s = rand.NewSource(time.Now().UnixNano())
		r = rand.New(s)
		sleep := r.Intn(150)
		time.Sleep(time.Millisecond * time.Duration(sleep))

		log.Printf("host=%s color=%s instanceCount=%d failAfter=%d sleep=%dms",
			host, color, instanceCount, failAfterCount, sleep)
	}
}

func ready(w http.ResponseWriter, r *http.Request) {
	log.Printf("ready()=%t", failed)
	if !failed {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
