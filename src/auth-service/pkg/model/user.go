package model

type DbUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"password"`
}
