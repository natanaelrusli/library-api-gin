package dto

type GetBookByIdParams struct {
	ID int `uri:"id" binding:"required"`
}
