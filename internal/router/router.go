package router

import (
	"net/http"

	"github.com/advaittrivedi1122/todolist/internal/handler"
)

// Returns instance of a new router
func NewRouter() http.Handler {
	router := http.NewServeMux()


	// API endpoints for HTTP server
	router.HandleFunc("/", handler.RootHandler)
	router.HandleFunc("/add-todo", handler.AddTodo)

	return router
}