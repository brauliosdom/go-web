package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
		w.Write([]byte("pong"))
	})

	if err := http.ListenAndServe("", nil); err != nil {
		panic(err)
	}
}
