package main

import (
	"html/template"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
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
}
