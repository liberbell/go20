package main

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

}
