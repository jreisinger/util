package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type router struct {
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprint(w, `
		<p><a href=addr>addr</a></p>
		<p><a href=headers>headers</a></p>
		`)
	case "/addr":
		addr := req.Header.Get("X-Forwarded-For") // behind proxy
		if addr == "" {
			addr = strings.Split(req.RemoteAddr, ":")[0]
		}
		fmt.Fprintf(w, "%v\n", addr)
	case "/headers":
		for name, values := range req.Header {
			fmt.Fprintf(w, "%v: %v\n", name, values)
		}
		return
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func main() {
	var r router
	// start a webserver
	log.Fatal(http.ListenAndServe(":5002", &r))
}
