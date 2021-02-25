package exception

type messageErr struct {
	ErrMessage    string `json:"message"`
	ErrHTTPStatus int    `json:"status"`
	ErrCode       string `json:"error"`
}

func (e *messageErr) Error() string {
	return e.ErrCode
}

func (e *messageErr) Text() string {
	return e.ErrMessage
}

func (e *messageErr) Status() int {
	return e.ErrHTTPStatus
}
