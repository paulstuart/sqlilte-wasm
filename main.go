package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// http.Handle("/", http.FileServer(http.Dir("site")))
	fs := middle(http.FileServer(http.Dir("site")))
	http.HandleFunc("/", fs)
	http.ListenAndServe(":"+port, nil)
}

func middle(hnd http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		hnd.ServeHTTP(w, r)
	}
}
