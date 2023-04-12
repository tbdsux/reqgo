package reqgo

import "encoding/json"

type Response struct {
	body []byte
}

func (res *Response) Text() string {
	return string(res.body)
}

func (res *Response) JSON(v interface{}) error {
	return json.Unmarshal(res.body, v)
}
