package types

type AddTodoRequest struct {
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type AddTodoResult struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}
