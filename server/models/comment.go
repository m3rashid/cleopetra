package models

type Comment struct {
	ID   string `validate:"required" json:"_id"`
	User *Auth  `validate:"required" json:"user"`
	Text string `validate:"required" json:"text"`
	Post *Post  `validate:"required" json:"post"`
}
