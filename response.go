package reqgo

import "encoding/json"

type Response struct {
	Body   []byte
	Code   int
	Status string
}

func (res *Response) Text() string {
	return string(res.Body)
}

func (res *Response) JSON(v interface{}) error {
	return json.Unmarshal(res.Body, v)
}
