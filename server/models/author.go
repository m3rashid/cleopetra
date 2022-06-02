package models

type Author struct {
	ID        string  `validate:"required" json:"_id"`
	User      *Auth   `validate:"required" json:"user"`
	Name      string  `validate:"required" json:"name"`
	Email     string  `validate:"required" json:"email"`
	Bio       string  `json:"bio"`
	Website   string  `json:"website"`
	Github    string  `json:"github"`
	Instagram string  `json:"instagram"`
	Twitter   string  `json:"twitter"`
	Photo     string  `json:"photo"`
	Featured  bool    `json:"featured"`
	Published bool    `json:"published"`
	Posts     []*Post `json:"posts"`
}
