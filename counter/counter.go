package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var globalCount int64

type hostData struct {
	Host          string
	Node          string
	LastSeen      int64
	Color         string
	HostCount     int64
	InstanceCount int64
}

var hosts = make(map[string]*hostData)
var started = time.Now()

func inc(w http.ResponseWriter, r *http.Request) {
	globalCount++

	vars := mux.Vars(r)
	color := vars["color"]
	host := vars["host"]
	node := vars["node"]
	ic := vars["instancecount"]
	instanceCount, _ := strconv.ParseInt(ic, 10, 64)

	data, ok := hosts[host]
	if ok == false {
		data = &hostData{host, node, 0, color, 0, 0}
		hosts[host] = data
	} else {
		data.Color = color
	}
	data.HostCount++
	data.LastSeen = time.Now().UnixNano() / int64(time.Millisecond)
	data.InstanceCount = instanceCount

	log.Printf("host=%s color=%s instanceCount=%d hostCount=%d globalCount=%d",
		data.Host, data.Color, instanceCount, data.HostCount, globalCount)
}

func get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	freshCount, staleCount := 0, 0
	showStale := vars["stale"]

	node := os.Getenv("NODE")
	if node == "" {
		node = "unknown"
	}

	fmt.Fprintln(w, "<html><head><meta http-equiv='refresh' content='3'></head><body style='font-size: 22px'><pre>")
	fmt.Fprintf(w, "Node: %s\n", node)
	fmt.Fprintf(w, "Started: %s\n", started.Format("2006-01-02 15:04:05"))
	fmt.Fprintf(w, "Current: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Fprintf(w, "Total Requests: %d\n", globalCount)

	table := "<table>"
	table += "<tr style='font-size: 22px'><td>&nbsp;</td><td>Container/Pod</td><td>Node</td><td align='right'>H</td><td align='right'>I</td></tr>"

	hostnames := make([]string, 0, len(hosts))
	for name := range hosts {
		hostnames = append(hostnames, name)
	}
	sort.Strings(hostnames)

	for _, h := range hostnames {
		v := hosts[h]

		color := v.Color
		threshold := time.Now().UnixNano() / int64(time.Millisecond)
		if v.LastSeen > threshold-1000 {
			freshCount++
		} else {
			staleCount++
			if showStale == "stale" {
				color = "grey"
			} else {
				continue
			}
		}

		table += fmt.Sprintf("<tr style='font-size: 22px'><td bgcolor='%s'>&nbsp;</td><td>%s&nbsp;</td><td>%s</td><td align='right'>&nbsp;%d</td><td align='right'>&nbsp;%d</td></tr>",
			color, v.Host, v.Node, v.HostCount, v.InstanceCount)
	}
	table += "</table>"
	fmt.Fprintf(w, "Fresh: %d  Stale: %d\n%s</pre></body></html>\n", freshCount, staleCount, table)
}

func clear(w http.ResponseWriter, r *http.Request) {
	for hostname, host := range hosts {
		threshold := time.Now().UnixNano() / int64(time.Millisecond)
		if host.LastSeen < threshold-10000 {
			delete(hosts, hostname)
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ready", ready)
	r.HandleFunc("/live", ready)
	go http.ListenAndServe(":8282", r)

	r = mux.NewRouter()
	r.HandleFunc("/get", get)
	r.HandleFunc("/get/{stale}", get)
	r.HandleFunc("/inc/{host}/{node}/{color}/{instancecount}", inc)
	r.HandleFunc("/clear", clear)

	http.ListenAndServe(":8181", r)
}

func ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
