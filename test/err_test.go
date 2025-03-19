package test

import (
	"errors"
	"fmt"
	"testing"
)

type MyError struct {
	Code int
}

func (e MyError) Error() string {
	return fmt.Sprintf("MyError: %d", e.Code)
}

var (
	ErrEmptyName  = errors.New("name is empty")
	ErrInvalidAge = errors.New("age is <= 0")
)

func TestXxx(t *testing.T) {
	// var err = MyError{Code: 123}
	// err2 := errors.New("err2")
	// t.Log(errors.New("err1") == errors.New("err1"))
	// t.Log(errors.As(err, &MyError{}))
	// wrapErr := fmt.Errorf("error on validation %w", err)
	// t.Log(errors.Unwrap(wrapErr))
}
