package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {

	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	srv.ListenAndServe()

}
