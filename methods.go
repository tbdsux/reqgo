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

// Post issues a POST request to the url. If `opt` is not nil,
// `opt.Data` will be set as request body and other options will
// be parsed to the request.
func Post(url string, opt *Options) (*Response, error) {
	return New(client).Post(url, opt)
}

// Post issues a POST request to the url. If `opt` is not nil,
// `opt.Data` will be set as request body and other options will
// be parsed to the request.
func (r *ReqGO) Post(url string, opt *Options) (*Response, error) {
	body, err := parseBody(opt)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	return r.doRequest(req, opt)
}

// Put issues a PUT request to the url. If `opt` is not nil,
// `opt.Data` will be set as request body and other options will
// be parsed to the request.
func Put(url string, opt *Options) (*Response, error) {
	return New(client).Put(url, opt)
}

// Put issues a PUT request to the url. If `opt` is not nil,
// `opt.Data` will be set as request body and other options will
// be parsed to the request.
func (r *ReqGO) Put(url string, opt *Options) (*Response, error) {
	body, err := parseBody(opt)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	return r.doRequest(req, opt)
}

// Patch issues a PATCH request to the url. If `opt` is not nil,
// `opt.Data` will be set as request body and other options will
// be parsed to the request.
func Patch(url string, opt *Options) (*Response, error) {
	return New(client).Patch(url, opt)
}

// Patch issues a PATCH request to the url. If `opt` is not nil,
// `opt.Data` will be set as request body and other options will
// be parsed to the request.
func (r *ReqGO) Patch(url string, opt *Options) (*Response, error) {
	body, err := parseBody(opt)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, url, body)
	if err != nil {
		return nil, err
	}

	return r.doRequest(req, opt)
}

// Delete issues a DELETE request to the url. If `opt` is not nil, it will be parsed to the request.
func Delete(url string, opt *Options) (*Response, error) {
	return New(client).Delete(url, opt)
}

// Delete issues a DELETE request to the url. If `opt` is not nil, it will be parsed to the request.
func (r *ReqGO) Delete(url string, opt *Options) (*Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	return r.doRequest(req, opt)
}

// Head issues a HEAD request to the url. If `opt` is not nil, it will be parsed to the request.
func Head(url string, opt *Options) (*Response, error) {
	return New(client).Head(url, opt)
}

// Head issues a HEAD request to the url. If `opt` is not nil, it will be parsed to the request.
func (r *ReqGO) Head(url string, opt *Options) (*Response, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}

	return r.doRequest(req, opt)
}
