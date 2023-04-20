package reqgo

import "testing"

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

func TestPut(t *testing.T) {
	r, err := Put("https://httpbin.org/put", &Options{
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
		t.Fatalf("put data seems not working")
	}
}

func TestPatch(t *testing.T) {
	r, err := Patch("https://httpbin.org/patch", &Options{
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
		t.Fatalf("patch data seems not working")
	}
}

func TestDelete(t *testing.T) {
	r, err := Delete("https://httpbin.org/delete", &Options{
		Headers: &Headers{
			"Accept": "application/json",
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if r.Code != 200 {
		t.Fatal("err: delete status code not 200")
	}
}

func TestHead(t *testing.T) {
	r, err := Head("https://httpbin.org", &Options{
		Headers: &Headers{
			"Accept": "application/json",
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if r.Code != 200 {
		t.Fatal("err: head status code not 200")
	}

}
