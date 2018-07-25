package main

import (
	"log"
	"net/http"
	"os"
)

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL)
		h.ServeHTTP(w, r)
	})
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	handler := logHandler(http.FileServer(http.Dir(path)))

	log.Println("Serving on http://localhost:5000")
	http.ListenAndServe(":5000", handler)
}
