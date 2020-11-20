package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", HelloServer)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	http.HandleFunc("/health_check", HealthCheck)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("received request!")
	if _, err := fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:]); err != nil {
		panic(err)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
