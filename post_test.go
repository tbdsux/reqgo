package reqgo

import (
	"testing"
)

func TestPost(t *testing.T) {
	r, err := Post("https://httpbin.org/post", &Options{
		Headers: &Headers{
			"Accept":       "application/json",
			"Content-Type": "application/json",
		},
		Data: map[string]string{
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

	if data["json"].(map[string]interface{})["hello"].(string) != "world" {
		t.Fatalf("post data seems not working")
	}
}
