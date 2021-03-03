package exception

import "encoding/json"

// NewAPIErrFromBytes creates a new api error from bytes
func NewAPIErrFromBytes(body []byte) (MessageErr, error) {
	var result messageErr
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
