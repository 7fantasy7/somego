package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"todoapi/internal/todo"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	todo.Init(router)

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		panic(err)
	}
}
