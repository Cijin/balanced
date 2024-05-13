package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type loadBalancer struct {
	backends []*url.URL
	current  int
}

func (lb *loadBalancer) nextBackend() *url.URL {
	// change to round robin
	return lb.backends[(lb.current+1)%len(lb.backends)]
}

func (lb *loadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend := lb.nextBackend()

	proxy := httputil.NewSingleHostReverseProxy(backend)
	proxy.ServeHTTP(w, r)
}

func main() {
	backends := []*url.URL{
		{Scheme: "http", Host: "localhost:8081"},
		{Scheme: "http", Host: "localhost:8082"},
	}

	go startServer("8081")
	go startServer("8082")

	lb := &loadBalancer{
		backends: backends,
		current:  0,
	}

	fmt.Println("load balancer listening on :8080")
	http.ListenAndServe(":8080", lb)
}
