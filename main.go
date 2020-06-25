package main

import (
	"log"
	"net/http"
)

type router struct {
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		root(w, req)
	case "/addr", "/ip", "/ipaddr":
		addr(w, req)
	case "/headers":
		headers(w, req)
	case "/status200":
		status200(w, req)
	case "/status302":
		status302(w, req)
	case "/status500":
		status500(w, req)
	default:
		http.Error(w, "404 Not Found", 404)
	}
}

func main() {
	var r router
	// start a webserver
	log.Fatal(http.ListenAndServe(":5002", &r))
}
