## Go Web Server with Concurrent Workers

### Overview
Go's implementation of web server already uses go routines on every request. So isn't it the best performance you can get ? Yes it is. But then what is the purpose of this project ?

This project provides a way for the web server to provide instant response but process the heavy work in the background using go routines. It is mostly useful, for resource heavy task to be done.

This repository contains a GoLang implementation of a web server that utilizes concurrent workers to handle incoming requests. The server responds to a "/hello" endpoint with an "Accepted" status code and the message "Welcome."

### Features

1. Concurrent Workers: The server employs a pool of worker goroutines (goroutines are lightweight threads in Go) to handle incoming requests concurrently. The number of worker goroutines is defined by the workers variable.

2. Request Handling: The "/hello" endpoint is defined in the handleHello function. When a request is received on this endpoint, the server responds with an "Accepted" status code and the message "Welcome." Additionally, a timestamp (in nanoseconds) is sent to a channel (ch1) to be processed by one of the worker goroutines.

3. Worker Functionality: The workerFxn function represents the behavior of the worker goroutines. Each worker waits for data from the ch1 channel, logs the received timestamp, and then simulates processing by sleeping for five seconds.


### Usage

1. Clone the repo 
```bash
git clone https://github.com/aswinbennyofficial/go-workerpool-webserver.git

```
```bash
cd go-workerpool-webserver
```

2. Start the server 
```bash
go run ./server/
```

3. Run the client to test it
```bash
go run ./client/
```
The server will start in port 8080. Can test it also using `http://localhost:8080/hello`
