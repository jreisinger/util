package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func root(w http.ResponseWriter, req *http.Request) {
	type page struct {
		Title        string
		WebUtilities []string
		CliUtilities map[string]string
	}
	p := page{
		Title:        "Utilities",
		WebUtilities: []string{"addr", "headers", "status200", "status500"},
		CliUtilities: map[string]string{
			"~/bin":      "https://github.com/jreisinger/dotfiles/tree/master/bin",
			"checkip":    "https://github.com/jreisinger/checkip",
			"runp":       "https://github.com/jreisinger/runp",
			"waf-runner": "https://github.com/jreisinger/waf-runner",
			"waf-tester": "https://github.com/jreisinger/waf-tester",
		},
	}
	t, err := template.New("page.html").ParseFiles("template/page.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func addr(w http.ResponseWriter, req *http.Request) {
	addr := req.Header.Get("X-Forwarded-For") // behind proxy
	if addr == "" {
		addr = strings.Split(req.RemoteAddr, ":")[0]
	}
	fmt.Fprintf(w, "%v\n", addr)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for _, name := range sortedKeys(req.Header) {
		values := strings.Join(req.Header[name], " | ")
		fmt.Fprintf(w, "%v: %v\n", name, values)
	}
}

func status200(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "200 OK")
}

func status500(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "500 Internal Server Error - a generic “catch-all” response", http.StatusInternalServerError)
}
