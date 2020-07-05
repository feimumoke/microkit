package util

type MyError struct {
	Code    int
	Message string
}

func (this *MyError) Error() string {
	return this.Message
}

func NewMyError(code int, msg string) error {
	return &MyError{Code: code, Message: msg,}
}
