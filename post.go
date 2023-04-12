package reqgo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// POST issues a POST request to the url. If `opt` is not nil,
// `opt.Data` will be set as request body and other options will
// be parsed to the request.
func Post(url string, opt *Options) (*Response, error) {
	return New(client).Post(url, opt)
}

// POST issues a POST request to the url. If `opt` is not nil,
// `opt.Data` will be set as request body and other options will
// be parsed to the request.
func (r *ReqGO) Post(url string, opt *Options) (*Response, error) {
	var data []byte = nil
	if opt != nil {
		j, err := json.Marshal(opt.Data)
		if err != nil {
			return nil, err
		}
		data = j
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	if opt != nil {
		// parse headers
		parseHeaders(req, opt.Headers)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Body:   body,
		Code:   res.StatusCode,
		Status: res.Status,
	}, nil
}
