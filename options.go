package reqgo

import "net/http"

type Headers map[string]string
type Params map[string]string

type Options struct {
	Headers *Headers
	Data    interface{}
	Params  *Params
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
