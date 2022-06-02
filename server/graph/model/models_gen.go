// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Auth struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type Author struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Photo     *string `json:"photo"`
	Bio       *string `json:"bio"`
	User      *Auth   `json:"user"`
	Featured  *bool   `json:"featured"`
	Published *bool   `json:"published"`
	Posts     []*Post `json:"posts"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Comment struct {
	ID      string `json:"id"`
	User    *Auth  `json:"user"`
	Comment string `json:"comment"`
}

type Like struct {
	ID   string `json:"id"`
	User *Auth  `json:"user"`
}

type Post struct {
	ID            string      `json:"id"`
	Title         string      `json:"title"`
	Slug          string      `json:"slug"`
	Excerpt       string      `json:"excerpt"`
	Content       string      `json:"content"`
	FeaturedImage *string     `json:"featuredImage"`
	Featured      bool        `json:"featured"`
	Published     bool        `json:"published"`
	Deleted       bool        `json:"deleted"`
	Categories    []*Category `json:"categories"`
	Author        *Auth       `json:"author"`
	Comments      []*Comment  `json:"comments"`
	Likes         []*Like     `json:"likes"`
}

type CreateAuthorInput struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignupInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type VerifyCreateAuthorInput struct {
	Email  string  `json:"email"`
	Photo  *string `json:"photo"`
	Bio    *string `json:"bio"`
	UserID string  `json:"userId"`
}

type Role string

const (
	RoleUser   Role = "USER"
	RoleAuthor Role = "AUTHOR"
	RoleEditor Role = "EDITOR"
	RoleAdmin  Role = "ADMIN"
)

var AllRole = []Role{
	RoleUser,
	RoleAuthor,
	RoleEditor,
	RoleAdmin,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleUser, RoleAuthor, RoleEditor, RoleAdmin:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}