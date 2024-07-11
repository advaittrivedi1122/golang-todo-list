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
	router.HandleFunc("/get-user-todo-by-id", handler.GetUserTodoById)
	router.HandleFunc("/get-user-todos", handler.GetUserTodos)
	router.HandleFunc("/update-user-todo-by-id", handler.UpdateUserTodoById)
	router.HandleFunc("/delete-user-todo-by-id", handler.DeleteUserTodoById)
	router.HandleFunc("/delete-user-todos", handler.DeleteUserTodos)

	return router
}