package main

import (
	"net/http"
)

func HandlerTLS() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", def)

	return mux
}

func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://"+addr+r.URL.Path, 307)
	})

	return mux
}

func def(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		infoLog.Println("404 not found")
		http.Error(w, "404 not found", 404)
		return
	}

	html, err := GetHTML("index.html")
	if err != nil {
		infoLog.Println("500 internal server error")
		http.Error(w, "500 internal server error", 500)
		return
	}

	w.WriteHeader(200)
	w.Write(html)
}
