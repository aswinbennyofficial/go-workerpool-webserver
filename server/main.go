package main

import (
	"log"
	"net/http"
	"time"
)

var (
	// ch1 is a channel to send request IDs to workers
	// Type of channel is int 
	// This channel can only send and receive integers
	// If you want to send and receive any data type, use interface{}
	ch1=make(chan int)
	// workers is the number of concurrent workers to spawn
	workers=50	
)

// workerFxn is a function that runs in a goroutine and processes
func workerFxn(){
	// range over the channel to receive request IDs
	for i:=range ch1{
		log.Println(i)
		// Simulate some work by sleeping for 5 seconds
		time.Sleep(time.Second * 5)
	}

}

func handleHello(w http.ResponseWriter, r * http.Request){
	// Send a response back to the client	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome"))
	
	// Send the request ID to the channel
	go func(){	
		// Sending the current time through the channel to process
		// You can send any data type through the channel as long as the
		// receiving end is expecting the same data type
		// In this case, we are sending the current time in nanoseconds

		ch1<-time.Now().Nanosecond()
	}()
}


func main(){

	// Spawn the workers
	for i:=0;i<workers;i++{
		go workerFxn()
	}

	// Register the handler function
	http.HandleFunc("/hello",handleHello)

	// Start the server
	log.Println("Starting server in 8080..")
	err:=http.ListenAndServe(":8080",nil)
	if err!=nil{
		log.Println("Error running on port 8080")
	}

}