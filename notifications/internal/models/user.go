package models

type User struct {
	Chat_ID  int64  `validate:"required"`
	Username string `validate:"required"`
	Name     string `validate:"required"`
}
