package model

type DbUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
