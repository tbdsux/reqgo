package reqgo

import (
	"io/ioutil"
	"net/http"
)

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
		// parse options
		parseHeaders(req, opt.Headers)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		body: body,
	}, nil
}
