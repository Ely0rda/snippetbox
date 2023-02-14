package main

import (
	"log"
	"net/http"
)

type homo struct{}

func (h *homo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello I am homo"))
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &homo{})

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
