package application

import (
	"net/http"

	"github.com/RalphTan37/microservice/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter() //creates new instance of chi router

	router.Use(middleware.Logger) //add middleware logger

	//http handler for / path
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	//creates an instance of the order routers
	orderHandler := &handler.Order{}

	//http methods
	router.Post("/", orderHandler.Create)
	router.Get("/", orderHandler.List)
	router.Get("/{id}", orderHandler.GetByID)
	router.Put("/{id}", orderHandler.UpdateByID)
	router.Delete("/{id}", orderHandler.DeleteByID)
}
