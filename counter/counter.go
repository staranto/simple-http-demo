package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var globalCount int64

type hostData struct {
	Host  string
	Color string
	Count int64
}

var hosts = make(map[string]*hostData)

func inc(w http.ResponseWriter, r *http.Request) {
	globalCount++

	vars := mux.Vars(r)
	color := vars["color"]
	host := vars["host"]
	instanceCount := vars["instancecount"]

	data, ok := hosts[host]
	if ok == false {
		data = &hostData{host, color, 0}
		hosts[host] = data
	} else {
		data.Color = color
	}
	data.Count++

	log.Printf("host=%s color=%s instanceCount=%s hostCount=%d globalCount=%d",
		data.Host, data.Color, instanceCount, data.Count, globalCount)
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<html><body style='font-size: 22px'><pre>")
	fmt.Fprintf(w, "Total Requests: %d", globalCount)

	fmt.Fprintln(w, "<table>")
	for _, v := range hosts {
		fmt.Fprintf(w, "<tr style='font-size: 22px'><td bgcolor='%s'>&nbsp;</td><td>%s</td><td>: %d</td></tr>",
			v.Color, v.Host, v.Count)
	}
	fmt.Fprintln(w, "</table></pre></body></html>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/get", get)
	r.HandleFunc("/inc/{host}/{color}/{instancecount}", inc)

	http.ListenAndServe(":8181", r)
}
