package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"
)

func root(w http.ResponseWriter, req *http.Request) {
	type page struct {
		Title        string
		WebUtilities []string
	}
	p := page{
		Title:        "Utilities",
		WebUtilities: []string{"ipaddr", "headers", "status200", "status302", "status500"},
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
	xff := req.Header["X-Forwarded-For"]
	if len(xff) > 0 { // behind proxy?
		fmt.Fprintf(w, "%s", strings.Split(xff[0], ", ")[0])
	} else {
		// Remove part with port number.
		port := regexp.MustCompile(`:\d+$`)
		ipAddr := port.ReplaceAllString(req.RemoteAddr, "")

		fmt.Fprintf(w, "%s", ipAddr)
	}
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

func status302(w http.ResponseWriter, req *http.Request) {
	scheme := "http" // proxy is handling TLS not us
	location := fmt.Sprintf("%s://%s/", scheme, req.Host)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusFound)
	fmt.Fprintln(w, "302 Found")
}

func status500(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "500 Internal Server Error - a generic “catch-all” response", http.StatusInternalServerError)
}

// favicon
func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon.ico")
}
