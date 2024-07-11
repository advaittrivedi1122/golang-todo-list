package types

import "time"

type AddTodoRequest struct {
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateUserTodoByIdRequest struct {
	UserId      int    `json:"user_id"`
	TodoId      int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type AddTodoResult struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

type UpdateUserTodoByIdResult = AddTodoResult

type DeleteUserTodoByIdResult = AddTodoResult

type DeleteUserTodosResult = AddTodoResult

type DeleteUserTodosRequest struct {
	UserId int `json:"user_id"`
}

type GetUserTodoByIdRequest struct {
	UserId int `json:"user_id"`
	TodoId int `json:"id"`
}

type DeleteUserTodoByIdRequest = GetUserTodoByIdRequest

type GetUserTodosRequest struct {
	UserId int    `json:"user_id"`
	Sort   string `json:"sort"`
	Filter string `json:"filter"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type GetUserTodosResponse struct {
	Todos []Todo `json:"todos"`
}

type Todo struct {
	UserId      int       `json:"user_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      string    `json:"status,omitempty"`
	TodoId      int       `json:"id,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	Updated     time.Time `json:"updated,omitempty"`
}

type ErrorResult = AddTodoResult
