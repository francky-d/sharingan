package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "text/html")

		resp.Write([]byte("<h1>Hello world</h1>"))
	})

	log.Fatal(http.ListenAndServe(":8000", router))
}
