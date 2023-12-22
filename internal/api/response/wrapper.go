package response

func NewSuccessResponseBody[T interface{}](payload T) ResponseWrapper[T] {
	return ResponseWrapper[T]{
		Data:    payload,
		Success: true,
		Errors:  []string{},
	}
}

type ResponseWrapper[T interface{}] struct {
	Data    T        `json:"data"`
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}
