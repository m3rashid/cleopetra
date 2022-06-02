package models

type Like struct {
	ID   string `validate:"required" json:"_id"`
	User *Auth  `validate:"required" json:"user"`
	Post *Post  `validate:"required" json:"post"`
}
