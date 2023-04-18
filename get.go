package reqgo

import (
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

	return r.doRequest(req, opt)
}
