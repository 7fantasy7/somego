package todo

import "github.com/go-chi/chi/v5"

func Init(router *chi.Mux) {
	router.Get("/todos", getTodos)
	router.Get("/todos/{id}", getTodo)
	router.Post("/todos", createTodo)
	router.Delete("/todos/{id}", deleteTodo)
	router.Put("/todos/{id}", updateTodo)
}
