package router

import (
	"github.com/ak1729/go-crud-mongodb/controller"
	"github.com/go-chi/chi/v5"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/movies", controller.GetMyAllMovies)
	r.Post("/api/movie", controller.CreateMovie)
	r.Put("/api/movie/{id}", controller.MarkAsWatched)
	r.Delete("/api/movie/{id}", controller.DeleteOneMovie)
	r.Delete("/api/deleteallmovie", controller.DeleteAllMovies)

	return r
}