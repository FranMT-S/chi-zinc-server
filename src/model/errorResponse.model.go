package model

import "fmt"

type ResponseError struct {
	Status int    `json:"Status"`
	Msg    string `json:"Msg"`
	Err    error  `json:"Err"`
}

// Return a Response Error
// type ResponseError struct {
// 	Status int    `json:"Status"`
// 	Msg    string `json:"Msg"`
// 	Err    error  `json:"Err"`
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
