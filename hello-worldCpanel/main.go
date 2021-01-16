package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello CPANEL... i'm Herya")
	})
	http.ListenAndServe("127.0.0.1:47000", nil)
}
