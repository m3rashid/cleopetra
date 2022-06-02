package models

type Post struct {
	ID            string      `validate:"required" json:"_id"`
	Author        *Author     `validate:"required" json:"author"`
	Title         string      `validate:"required" json:"title"`
	Slug          string      `validate:"required" json:"slug"`
	Excerpt       string      `validate:"required" json:"excerpt"`
	Content       string      `validate:"required" json:"text"`
	MinutesToRead int         `validate:"required" json:"minutesToRead"`
	FeaturedImage string      `validate:"required" json:"featuredImage"`
	Featured      bool        `validate:"required" json:"featured"`
	Published     bool        `validate:"required" json:"published"`
	Deleted       bool        `validate:"required" json:"deleted"`
	Categories    []*Category `validate:"required" json:"categories"`
	Comments      []*Comment  `json:"comments"`
	Likes         []*Like     `json:"likes"`
	Tags          []string    `json:"tags"`
}
