package dto

type CreateOneBookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Cover       string `json:"cover" binding:"required"`
	AuthorId    int32  `json:"author_id" binding:"required"`
	Stock       int32  `json:"stock" binding:"required"`
}
