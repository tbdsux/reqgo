package reqgo

import "encoding/json"

type Response struct {
	Body   []byte
	Code   int
	Status string
}

// Parse response body to string.
func (res *Response) Text() string {
	return string(res.Body)
}

// Parse response body as JSON to v.
func (res *Response) JSON(v interface{}) error {
	return json.Unmarshal(res.Body, v)
}
