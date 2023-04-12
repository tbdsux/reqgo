package reqgo

import "net/http"

type Headers map[string]string

type Options struct {
	Headers Headers
	Data    interface{}
}

func parseHeaders(req *http.Request, headers Headers) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
