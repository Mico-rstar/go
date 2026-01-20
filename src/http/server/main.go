package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping/{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("req: ping")
		w.Write([]byte("pong"))
	})
	http.ListenAndServe(":8080", nil)
}