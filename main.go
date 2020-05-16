package main

import (
	"log"
	"net/http"
	"sort"
)

type router struct {
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		root(w, req)
	case "/addr":
		addr(w, req)
	case "/headers":
		headers(w, req)
	case "/status/200":
		status200(w, req)
	case "/status/500":
		status500(w, req)
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
