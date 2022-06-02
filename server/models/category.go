package models

type Category struct {
	ID   string `validate:"required" json:"_id"`
	Name string `validate:"required" json:"name"`
	Slug string `validate:"required" json:"slug"`
}
