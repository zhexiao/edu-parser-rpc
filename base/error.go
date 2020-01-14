package base

type EpError struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func NewEpError(code int, msg string) *EpError {
	return &EpError{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}

func (e *EpError) Error() string {
	return e.ErrorMsg
}
