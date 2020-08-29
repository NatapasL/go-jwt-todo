package forms

type CreateTodoParams struct {
	UserID uint   `json:"user_id"`
	Title  string `json:"title"`
}
