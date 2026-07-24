package result

import "fmt"

type Result[T any] struct {
	value T
	err error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{value: value}
}

func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(fmt.Sprintf("result: called Unwrap on Err: %v", r.err))
	}
	return r.value
}

func (r Result[T]) UnwrapOr(fallback T) T {
	if r.err != nil {
		return fallback
	}
	return r.value
}

func (r Result[T]) Error() error {
	return r.err
}

func (r Result[T]) Get() (T, error) {
	return r.value, r.err
}

func Map[T, U any](r Result[T], fn func(T) U) Result[U] {
	if r.err != nil {
		return Err[U](r.err)
	}
	return Ok(fn(r.value))
}