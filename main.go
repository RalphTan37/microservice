package main

import (
	"context"
	"fmt" //format package

	"github.com/RalphTan37/microservice/application"
)

func main() {
	app := application.New() //new instance of the app

	err := app.Start(context.TODO()) //start the app
	//handles error
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
