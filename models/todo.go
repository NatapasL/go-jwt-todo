package models

type TodoModel struct {
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}
