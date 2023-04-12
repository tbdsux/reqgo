package reqgo

import "net/http"

var client = http.DefaultClient

type ReqGO struct {
	client *http.Client
}

func New(cl *http.Client) *ReqGO {
	return &ReqGO{
		client: cl,
	}
}
