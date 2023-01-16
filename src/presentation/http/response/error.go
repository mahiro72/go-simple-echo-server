package response

type ErrResponse struct {
	Message string `json:"message"`
}

func NewErr(err error) *ErrResponse {
	return &ErrResponse{
		Message: err.Error(),
	}
}
