package application

import (
	"context" //manage lifetime and cancellation of tasks
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

// store any app dependencies
type App struct {
	router http.Handler
	rdb    *redis.Client //store redis client
}

// constructor for app
func New() *App {
	app := &App{
		router: loadRoutes(),
		rdb:    redis.NewClient(&redis.Options{}), //creates a new instance of redis
	}
	return app
}

// func def to start app
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{ //define server
		Addr:    ":3000",  //same addr
		Handler: a.router, //handler to app router
	}

	err := a.rdb.Ping(ctx).Err() //calls the Ping method of the redit client
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	//anonymous function
	defer func() {
		//error check
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("Starting server")

	ch := make(chan error, 1) //create channel error type, buffer sz 1

	//goroutine - run server concurrently
	//starts a new anonymous function in a new thread of execution
	//ensures that it doesn't block the main thread
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err) //publish value onto channel
		}
		close(ch) //closes channel when function is done
	}()

	select {
	//receiver for channel
	//blocks the code execution until it either receives a value or the channel is closed
	case err = <-ch: //capture any value sent on this channel
		return err
	case <-ctx.Done(): //returns channel inside
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		return server.Shutdown(timeout)
	}

	return nil //if everything works okay
}
