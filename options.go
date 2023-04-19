package reqgo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Headers map[string]string
type Params map[string]string

type Options struct {
	Headers *Headers
	Data    interface{}
	Params  *Params
}

// prase opt.Data if not nil
func parseBody(opt *Options) (io.Reader, error) {
	if opt != nil && opt.Data != nil {
		data, err := json.Marshal(opt.Data)
		if err != nil {
			return nil, err
		}

		return bytes.NewBuffer(data), nil
	}

	return nil, nil
}

// parse headers to the request
func parseHeaders(req *http.Request, headers *Headers) {
	if headers == nil {
		return
	}

	for key, value := range *headers {
		req.Header.Set(key, value)
	}
}

// parse query params to the request
func parseParams(req *http.Request, params *Params) {
	if params == nil {
		return
	}

	q := req.URL.Query()

	for key, value := range *params {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()
}
