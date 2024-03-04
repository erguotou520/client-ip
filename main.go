package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		fmt.Fprintf(w, "client ip: %s\n", ip)
	})

	fmt.Printf("Server started on port 8080\n")
	http.ListenAndServe(":8080", nil)
}
