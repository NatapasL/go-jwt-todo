package forms

type CreateTodoParams struct {
	UserID uint64 `json:"user_id"`
	Title  string `json:"title"`
}
