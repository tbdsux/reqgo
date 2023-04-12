package reqgo

import (
	"io"
	"net/http"
)

// Get issues a GET request to the url. If `opt` is not nil, it will be parsed to the request.
func Get(url string, opt *Options) (*Response, error) {
	return New(client).Get(url, opt)
}

// Get issues a GET request to the url. If `opt` is not nil, it will be parsed to the request.
func (r *ReqGO) Get(url string, opt *Options) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
