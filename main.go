package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"text/template"
)

type router struct {
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		type page struct {
			Title        string
			WebUtilities []string
			CliUtilities map[string]string
		}
		p := page{
			Title:        "Utilities",
			WebUtilities: []string{"addr", "headers"},
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
