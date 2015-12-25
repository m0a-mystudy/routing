package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/k0kubun/pp"
)

// func main() {
// 	http.ListenAndServe(":3001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Helloworld!")
// 	}))
// }

func stripPrefix(prefix string, h http.Handler) http.Handler {
	if prefix == "" {
		return h
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
			r.URL.Path = p
			pp.Println(p)
			h.ServeHTTP(w, r)
		} else {
			http.NotFound(w, r)
		}
	})
}

func main() {

	apiMux := http.NewServeMux()
	apiMux.Handle("/api1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "in /api/api1")
	}))
	apiMux.Handle("/api2", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "in /api/api2")
	}))

	http.Handle("/api/v1/", stripPrefix("/api/v1", apiMux))
	http.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello!")
	}))
	// http.Handle("/", http.RedirectHandler("/hello", 301))

	http.ListenAndServe(":3001", nil)

}
