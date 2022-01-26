package web

type CategoryUpdateRequest struct {
	Id           int    `validate:"required"`
	CategoryName string `validate:"required,min=1,max=100" json:"category_name"`
	CategorySlug string `validate:"required,min=1,max=200" json:"category_slug"`
}
