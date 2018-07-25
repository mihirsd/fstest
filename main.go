package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Serving on http://localhost:5000")
	http.ListenAndServe(":5000", http.FileServer(http.Dir(path)))
}
