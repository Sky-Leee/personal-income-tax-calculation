package apis

type TaxRequest struct {
	Income float64 `json:"income" binding:"required"`
}

type TaxResponse struct {
	Message string `json:"msg"`
	Tax float64 `json:"tax"`
}