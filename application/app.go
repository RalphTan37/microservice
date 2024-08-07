package application

import (
	"context" //manage lifetime and cancellation of tasks
	"fmt"
	"net/http"
)

// store any app dependencies
type App struct {
	router http.Handler
}

// constructor for app
func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app
}

// func def to start app
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{ //define server
		Addr:    ":3000",  //same addr
		Handler: a.router, //handler to app router
	}

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err) //error wrapping - wrap error w/ another error
	}

	return nil //if everything works okay
}
