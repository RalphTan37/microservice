package main

import (
	"context"
	"fmt" //format package
	"os"
	"os/signal" //access to incoming signals

	"github.com/RalphTan37/microservice/application"
)

func main() {
	app := application.New() //new instance of the app

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt) //takes in a signal and a context, return another context that will be notified if the signal is created
	//context parameter derives a new context from

	defer cancel() //signals to go that the cancel function should be called at the end of the current functionn it's in

	err := app.Start(ctx) //responds to any interrupt signals
	//handles error
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
