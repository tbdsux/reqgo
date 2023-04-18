package reqgo

import (
	"io"
	"net/http"
)

var client = http.DefaultClient

type ReqGO struct {
	client *http.Client
}

func New(cl *http.Client) *ReqGO {
	return &ReqGO{
		client: cl,
	}
}

func (r *ReqGO) doRequest(req *http.Request, opt *Options) (*Response, error) {
	if opt != nil {
		if opt.Headers != nil {
			// parse headers
			parseHeaders(req, opt.Headers)
		}

		if opt.Params != nil {
			// parse params
			parseParams(req, opt.Params)
		}
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
