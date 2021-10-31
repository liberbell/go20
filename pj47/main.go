package main

import "fmt"

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "Custom Error happend.", ErrCode: 11}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	// fmt.Println(err.Messageß)*
	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}
}
