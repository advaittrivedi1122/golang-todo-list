package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/advaittrivedi1122/todolist/database"
	"github.com/advaittrivedi1122/todolist/internal/types"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Todos List!!"))
}

func AddTodo(w http.ResponseWriter, r *http.Request)  {
	var req types.AddTodoRequest
	var res types.AddTodoResult
	json.NewDecoder(r.Body).Decode(&req)
	w.Header().Add("Content-type", "application/json")

	if err := database.InsertUserTodo(req); err != nil {
		log.Printf("Unable to Insert User Todo : %v", err)
		res = types.AddTodoResult{
			Success: false,
			Error: err.Error(),
		}
		resBytes, _ := json.Marshal(res)
		w.Write(resBytes)
		return
	}
	
	if err := database.IncrementUserTodosCount(req.UserId); err != nil {
		log.Printf("Unable to increment user todos count : %v", err)
		res = types.AddTodoResult{
			Success: false,
			Error: err.Error(),
		}
		resBytes, _ := json.Marshal(res)
		w.Write(resBytes)
		return
	}

	res = types.AddTodoResult{
		Success: true,
	}
	resBytes, _ := json.Marshal(res)
	w.Write(resBytes)
}