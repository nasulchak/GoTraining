package main

import (
	"errors"
	"fmt"
)

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Operation completed successfuly!")
}

type customError struct {
	code    int
	message string
	err     error
}

// Error returs the error message. Implementing Error() method of error interface
func (e *customError) Error() string {
	return fmt.Sprintf("Error %d: %s, %v", e.code, e.message, e.err)
}

func (e *customError) Unwrap() error {
	return e.err
}

// Function than return a custom error
//
//	func doSomething() error {
//		return &customError{
//			code:    500,
//			message: "Something went wrong!",
//		}
//	}

var errInternal = errors.New("internal error")

func doSomething() error {
	// err := doSomethingElse()
	// if err != nil {
	// 	return &customError{
	// 		code:    500,
	// 		message: "Something went wrong",
	// 		err:      err,
	// 	}
	// }
	err := &customError{
		code:    500,
		message: "Something went wrong",
		err:     errInternal,
	}

	if errors.Is(err, errInternal) {
		fmt.Println("This is an internal error")
	}

	var ce *customError
	if errors.As(err, &ce) {
		fmt.Println("Code error:", ce.code)
	}

	return nil
}

func doSomethingElse() error {
	return errors.New("internal error")
}
