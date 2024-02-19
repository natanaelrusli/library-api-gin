package httpdto

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type UpdateStockResponse struct {
	Id    int `json:"id"`
	Stock int `json:"stock"`
}
