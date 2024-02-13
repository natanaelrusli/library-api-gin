package dto

type Query struct {
	Name string `form:"name"`
}

type GetBookByIdParams struct {
	ID int `uri:"id" binding:"required"`
}

type GetAuthorByBookIdParams struct {
	ID int `uri:"id" binding:"required"`
}
