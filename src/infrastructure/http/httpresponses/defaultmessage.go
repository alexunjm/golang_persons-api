package httpresponses

type defaultMessage struct {
	Message    string      `json:"message"`
	HTTPStatus int         `json:"status"`
	Code       string      `json:"code"`
	Data       interface{} `json:"data"`
}

func (m *defaultMessage) StringCode() string {
	return m.Code
}

func (m *defaultMessage) Text() string {
	return m.Message
}

func (m *defaultMessage) Status() int {
	return m.HTTPStatus
}
