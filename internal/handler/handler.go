package handler

import (
	"encoding/json"
	"fmt"
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

func GetUserTodoById(w http.ResponseWriter, r *http.Request) {
	var req types.GetUserTodoByIdRequest
	json.NewDecoder(r.Body).Decode(&req)
	w.Header().Add("Content-type", "application/json")

	todo, err := database.GetUserTodoById(req.UserId, req.TodoId);
	if err != nil {
		log.Printf("Unable to Get User Todo : %v", err)
		errorResult := types.ErrorResult{
			Success: false,
			Error: err.Error(),
		}
		resBytes, _ := json.Marshal(errorResult)
		w.Write(resBytes)
		return
	}

	resBytes, _ := json.Marshal(todo)
	w.Write(resBytes)
}

func GetUserTodos(w http.ResponseWriter, r *http.Request) {
	var req types.GetUserTodosRequest
	var res types.GetUserTodosResponse
	json.NewDecoder(r.Body).Decode(&req)
	w.Header().Add("Content-type", "application/json")

	fmt.Printf("\n\n[Req] : %+v\n\n", req)

	if (req.Filter != "pending" && req.Filter != "completed") {
		req.Filter = ""
	}

	if (req.Sort != "desc") {
		req.Sort = ""
	}

	todos, err := database.GetUserTodos(req);
	if err != nil {
		log.Printf("Unable to Get User Todos : %v", err)
		errorResult := types.ErrorResult{
			Success: false,
			Error: err.Error(),
		}
		resBytes, _ := json.Marshal(errorResult)
		w.Write(resBytes)
		return
	}

	res = types.GetUserTodosResponse{
		Todos: todos,
	}

	resBytes, _ := json.Marshal(res)
	w.Write(resBytes)
}

func UpdateUserTodoById(w http.ResponseWriter, r *http.Request) {
	var req types.UpdateUserTodoByIdRequest
	var res types.UpdateUserTodoByIdResult
	json.NewDecoder(r.Body).Decode(&req)
	w.Header().Add("Content-type", "application/json")

	if (req.UserId <= 0) {
		res = types.UpdateUserTodoByIdResult{
			Success: false,
			Error: "user_id missing",
		}
		resBytes, _ := json.Marshal(res)
		w.Write(resBytes)
		return
	}
	if (req.TodoId <= 0) {
		res = types.UpdateUserTodoByIdResult{
			Success: false,
			Error: "id missing",
		}
		resBytes, _ := json.Marshal(res)
		w.Write(resBytes)
		return
	}

	if err := database.UpdateUserTodoById(req); err != nil {
		log.Printf("Unable to Update User Todo : %v", err)
		res = types.UpdateUserTodoByIdResult{
			Success: false,
			Error: err.Error(),
		}
		resBytes, _ := json.Marshal(res)
		w.Write(resBytes)
		return
	}

	res = types.UpdateUserTodoByIdResult{
		Success: true,
	}
	resBytes, _ := json.Marshal(res)
	w.Write(resBytes)
}

func DeleteUserTodoById(w http.ResponseWriter, r *http.Request) {
	var req types.DeleteUserTodoByIdRequest
	var res types.DeleteUserTodoByIdResult
	json.NewDecoder(r.Body).Decode(&req)
	w.Header().Add("Content-type", "application/json")

	if (req.UserId <= 0) {
		res = types.DeleteUserTodoByIdResult{
			Success: false,
			Error: "user_id missing",
		}
		resBytes, _ := json.Marshal(res)
		w.Write(resBytes)
		return
	}
	if (req.TodoId <= 0) {
		res = types.DeleteUserTodoByIdResult{
			Success: false,
			Error: "id missing",
		}
		resBytes, _ := json.Marshal(res)
		w.Write(resBytes)
		return
	}

	if err := database.DeleteUserTodoById(req); err != nil {
		log.Printf("Unable to Delete User Todo : %v", err)
		res = types.DeleteUserTodoByIdResult{
			Success: false,
			Error: err.Error(),
		}
		resBytes, _ := json.Marshal(res)
		w.Write(resBytes)
		return
	}

	if err := database.DecrementUserTodosCount(req.UserId); err != nil {
		log.Printf("Unable to decrement user todos count : %v", err)
		res = types.DeleteUserTodoByIdResult{
			Success: false,
			Error: err.Error(),
		}
		resBytes, _ := json.Marshal(res)
		w.Write(resBytes)
		return
	}

	res = types.DeleteUserTodoByIdResult{
		Success: true,
	}
	resBytes, _ := json.Marshal(res)
	w.Write(resBytes)
}

func DeleteUserTodos(w http.ResponseWriter, r *http.Request) {
	var req types.DeleteUserTodosRequest
	var res types.DeleteUserTodosResult
	json.NewDecoder(r.Body).Decode(&req)
	w.Header().Add("Content-type", "application/json")

	if (req.UserId <= 0) {
		res = types.DeleteUserTodosResult{
			Success: false,
			Error: "user_id missing",
		}
		resBytes, _ := json.Marshal(res)
		w.Write(resBytes)
		return
	}

	if err := database.DeleteUserTodos(req.UserId); err != nil {
		log.Printf("Unable to Delete User Todos : %v", err)
		res = types.DeleteUserTodosResult{
			Success: false,
			Error: err.Error(),
		}
		resBytes, _ := json.Marshal(res)
		w.Write(resBytes)
		return
	}

	if err := database.ResetUserTodosCount(req.UserId); err != nil {
		log.Printf("Unable to Delete User Todos : %v", err)
		res = types.DeleteUserTodosResult{
			Success: false,
			Error: err.Error(),
		}
		resBytes, _ := json.Marshal(res)
		w.Write(resBytes)
		return
	}

	res = types.DeleteUserTodosResult{
		Success: true,
	}
	resBytes, _ := json.Marshal(res)
	w.Write(resBytes)
}