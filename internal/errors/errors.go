package errors

type Error struct {
	Msg        string `json:"msg"`
	Detail     string `json:"detail"`
	HTTPStatus int    `json:"-"`
	Err        error  `json:"-"`
}

func NewError() *Error {
	return &Error{}
}
func (e *Error) NewError(msg, detail string, httpStatus int, err error) *Error {
	return &Error{
		Msg:        msg,
		Detail:     detail,
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
