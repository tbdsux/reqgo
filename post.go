package reqgo

import (
	"bytes"
	"encoding/json"
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

	return r.doRequest(req, opt)
}
