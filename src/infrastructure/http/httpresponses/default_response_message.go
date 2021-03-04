package httpresponses

type defaultResponseMessage struct {
	Info defaultMessage `json:"info"`
	Data interface{}    `json:"data"`
}

func (m *defaultResponseMessage) StringCode() string {
	return m.Info.Code
}

func (m *defaultResponseMessage) Text() string {
	return m.Info.Message
}

func (m *defaultResponseMessage) Status() int {
	return m.Info.HTTPStatus
}
