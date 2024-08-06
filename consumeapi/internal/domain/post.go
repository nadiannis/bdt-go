package domain

type Post struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
