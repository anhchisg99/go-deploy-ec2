package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("thang"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	log.Print("Starting server on :4001")
	err := http.ListenAndServe(":4001", mux)
	log.Fatal(err)
}
