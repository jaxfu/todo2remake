package types

type Todo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	ID      uint   `json:"id"`
}
