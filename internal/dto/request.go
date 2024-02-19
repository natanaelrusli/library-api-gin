package dto

type CreateOneBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Cover       string `json:"cover" binding:"required"`
	AuthorId    int32  `json:"author_id" binding:"required"`
	Stock       int32  `json:"stock" binding:"required"`
}

type CreateBorrowingRecordRequest struct {
	UserId int    `json:"user_id" binding:"required"`
	BookId int    `json:"book_id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type UpdateBookStockRequest struct {
	BookId int `json:"book_id" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}
