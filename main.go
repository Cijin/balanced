package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var ports = []string{
	"8081", "8082", "8083", "8084", "8085",
}

type loadBalancer struct {
	backends []*url.URL
	current  int
}

func (lb *loadBalancer) nextBackend() *url.URL {
	// change to round robin
	be := lb.backends[lb.current]
	lb.current = (lb.current + 1) % len(lb.backends)

	return be
}

func (lb *loadBalancer) errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Println("Error fulfilling request on server idx:", lb.current)

	lb.ServeHTTP(w, r)
}

func (lb *loadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend := lb.nextBackend()

	proxy := httputil.NewSingleHostReverseProxy(backend)
	proxy.ErrorHandler = lb.errorHandler
	proxy.ServeHTTP(w, r)
}

func main() {
	var backends []*url.URL
	for _, port := range ports {
		backends = append(backends, &url.URL{Scheme: "http", Host: fmt.Sprintf("localhost:%s", port)})

		// TODO:create cancellable context, that gets called when load balancer quits
		go startServer(port)
	}

	lb := &loadBalancer{
		backends: backends,
		current:  0,
	}

	fmt.Println("load balancer listening on :8080")
	http.ListenAndServe(":8080", lb)
}
