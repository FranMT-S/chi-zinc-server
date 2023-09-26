package model

import "fmt"

type ResponseError struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Err    error  `json:"error"`
}

// Return a Response Error
// type ResponseError struct {
// 	Status int    `json:"status"`
// 	Msg    string `json:"msg"`
// 	Err    error  `json:"error"`
// }
func NewResponseError(status int, msg string, err error) *ResponseError {
	return &ResponseError{status, msg, err}
}

func (err *ResponseError) Error() string {
	return fmt.Sprintf(`{
			"status":%v,
			"msg":"%v",
			"error":"%v"
		}`, err.Status, err.Msg, err.Err)
}
