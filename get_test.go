package reqgo

import (
	"testing"
)

func TestGet(t *testing.T) {
	r, err := Get("https://httpbin.org/get", &Options{
		Headers: &Headers{
			"Accept": "application/json",
		},
		Params: &Params{
			"hello": "world",
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	var data map[string]interface{}
	if err := r.JSON(&data); err != nil {
		t.Fatal(err)
	}

	if data["args"].(map[string]interface{})["hello"].(string) != "world" {
		t.Fatalf("error: request params seems to be not working")
	}
}
