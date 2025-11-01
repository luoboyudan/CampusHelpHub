package errors

import "fmt"

type Error struct {
	Msg        string `json:"msg"`
	Detail     string `json:"detail"`
	HTTPStatus int    `json:"-"`
	Err        error  `json:"-"`
}

func NewError() *Error {
	return &Error{}
}

// NewError 创建一个新的错误
func (e *Error) NewError(errType string, httpStatus int, err error) *Error {
	if err == nil {
		err = fmt.Errorf("%s", errType)
	}
	return &Error{
		Msg:        ErrorMsgTemplates[errType].Msg,
		Detail:     fmt.Sprintf(ErrorMsgTemplates[errType].Detail, err.Error()),
		HTTPStatus: httpStatus,
		Err:        err,
	}
}
func (e *Error) Error() string {
	return e.Msg + ": " + e.Detail
}

func (e *Error) GetHTTPStatus() int {
	return e.HTTPStatus
}

func (e *Error) GetErr() error {
	return e.Err
}

func (e *Error) Is(target error) bool {
	return e.Err == target
}

func (e *Error) GetDetail() string {
	return e.Detail
}
