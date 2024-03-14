package service

type IResponse[T any] interface {
	Result() T
	HasError() bool
	AddResult(value T)
	AddError(err error)
}

type Response[T any] struct {
	error error
	value T
}

func (r *Response[T]) HasError() bool {
	return r.error != nil
}

func (r *Response[T]) Result() T {
	return r.value
}

func (r *Response[T]) AddResult(value T) {
	r.value = value
}

func (r *Response[T]) AddError(err error) {
	r.error = err
}
