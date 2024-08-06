package main

import (
	"fmt"
	"net/http" //functionality for creating http clients & servers
)

func basicHandler(w http.ResponseWriter, r *http.Request) { //write & request parameters
	w.Write([]byte("Hello World!"))
}

func main() {
	server := &http.Server{
		Addr:    ":3000",                        //server address
		Handler: http.HandlerFunc(basicHandler), //interface when the server receives a request
	}

	err := server.ListenAndServe() //initializes and starts an HTTP server; if the server encounters an error, it is captured in the err var
	//checks if the error exists
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}
