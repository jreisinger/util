package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
)

type router struct {
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		root(w, req)
	case "/addr":
		addr := req.Header.Get("X-Forwarded-For") // behind proxy
		if addr == "" {
			addr = strings.Split(req.RemoteAddr, ":")[0]
		}
		fmt.Fprintf(w, "%v\n", addr)
	case "/headers":
		for _, name := range sortedKeys(req.Header) {
			values := strings.Join(req.Header[name], " | ")
			fmt.Fprintf(w, "%v: %v\n", name, values)
		}
	case "/status/200":
		fmt.Fprintln(w, "200 OK")
	case "/status/500":
		http.Error(w, "500 Internal Server Error - a generic “catch-all” response", http.StatusInternalServerError)
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

// return keys of a map alphabetically sorted
func sortedKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func main() {
	var r router
	// start a webserver
	log.Fatal(http.ListenAndServe(":5002", &r))
}
