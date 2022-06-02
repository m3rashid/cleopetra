package models

type Auth struct {
	ID       string `validate:"required" json:"_id"`
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Role     string `validate:"required" json:"role"`
}
