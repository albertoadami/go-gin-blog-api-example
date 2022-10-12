package json

type GenericErrorResponse struct {
	Error       string `json:"error" binding:"required"`
	Description string `json:"description" binding:"required"`
}
