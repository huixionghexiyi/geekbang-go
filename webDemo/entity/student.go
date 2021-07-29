package entity

type Student struct {
	Id       int `json:"id"`
	Username string `json:"username"`
	Age      int `json:"age"`
	Gender   int `json:"gender"`
	Email    string `json:"email"`
}
