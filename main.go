package main

import (
	"fmt"      //format package
	"net/http" //functionality for creating http clients & servers

	"github.com/go-chi/chi"            //chi package
	"github.com/go-chi/chi/middleware" //logging middleware
)

// handler for an http server
func basicHandler(w http.ResponseWriter, r *http.Request) { //write & request parameters
	w.Write([]byte("Hello World!"))
}

func main() {
	router := chi.NewRouter()     //initializes router
	router.Use(middleware.Logger) //add middleware to router

	router.Get("/Hello", basicHandler) //defines route for HTTP get request

	//instantiates an http server
	server := &http.Server{
		Addr:    ":3000", //server address
		Handler: router,  //interface when the server receives a request
	}

	err := server.ListenAndServe() //initializes and starts an HTTP server; if the server encounters an error, it is captured in the err var
	//checks if the error exists
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}
