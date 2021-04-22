package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	log.Println("Listening at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Write([]byte("Hello World"))
	w.WriteHeader(http.StatusOK)
}
