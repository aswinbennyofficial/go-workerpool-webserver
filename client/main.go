package main

import (
	"fmt"
	"net/http"
	"sync"
)

const (
	// url is the URL of the server to send requests to
	url = "http://localhost:8080/hello"
	// totalRequests is the number of requests to send
	totalRequests = 10000
)

func main() {
	// Create a WaitGroup to wait for all requests to finish
	var wg sync.WaitGroup

	// Send totalRequests number of requests
	for i := 1; i <= totalRequests; i++ {
		wg.Add(1)
		go sendRequest(i, &wg)
	}

	// Wait for all requests to finish
	wg.Wait()

	fmt.Println("All requests completed.")
}

func sendRequest(requestID int, wg *sync.WaitGroup) {
	// Defer the call to Done to mark this function as finished
	defer wg.Done()

	// Send a GET request to the server
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Request %d failed: %v\n", requestID, err)
		return
	}
	
	fmt.Printf("Request %d completed %s\n", requestID,resp.Status)
}
