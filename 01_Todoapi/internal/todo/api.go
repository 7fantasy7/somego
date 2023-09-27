package todo

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var (
	todos = []Todo{{ID: 1, Title: "Take kid from school"}, Todo{ID: 2, Title: "Buy bottle of wine"}}
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		return
	}
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, item := range todos {
		if item.ID == id {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				break
			}
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	for index, item := range todos {
		if item.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var newTodo Todo
	_ = json.NewDecoder(r.Body).Decode(&newTodo)
	for index, item := range todos {
		if item.ID == id {
			todos[index] = Todo{ID: id, Title: newTodo.Title}
		}
	}
	err := json.NewEncoder(w).Encode(todos)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
}
