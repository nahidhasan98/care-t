package model

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Gender  string `json:"gender"`
	Status  string `json:"status"`
	Message string `json:"message"`
	Field   string `json:"field"`
}
