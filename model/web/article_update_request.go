package web

type ArticleUpdateRequest struct {
	Id         int    `validate:"required"`
	Title      string `validate:"required,min=1,max=100" json:"title"`
	Slug       string `validate:"required,min=1,max=50" json:"slug"`
	CategoryId int    `validate:"required" json:"category_id"`
	Content    string `validate:"required,min=10" json:"content"`
}
