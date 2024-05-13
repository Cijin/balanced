package main

import (
	"fmt"
	"log"
	"net/http"
)

func startServer(port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var headers string
		for k, v := range r.Header {
			headers += fmt.Sprintf("%s:%s\n", k, v)
		}
		log.Println(headers)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello from the backend:%s\n", port)))
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Println("backend listening on port:", port)
	log.Fatal(server.ListenAndServe())
}
