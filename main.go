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

	http.ListenAndServe(":8080", nil)
}
