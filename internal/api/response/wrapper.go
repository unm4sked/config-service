package response

func NewSuccessResponseBody[T interface{}](payload T) ResponseWrapper[T] {
	return ResponseWrapper[T]{
		Data:    payload,
		Success: true,
		Errors:  []string{},
	}
}

func NewFailureResponseBody(errors []string) ResponseWrapper[any] {
	return ResponseWrapper[any]{
		Data:    nil,
		Success: true,
		Errors:  errors,
	}
}

type ResponseWrapper[T interface{}] struct {
	Data    T        `json:"data"`
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}
