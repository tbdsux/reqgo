package main

import (
	"fmt"
	"log"

	"github.com/tbdsux/reqgo"
)

func main() {
	r, err := reqgo.Get("https://httpbin.org/get", &reqgo.Options{
		Headers: reqgo.Headers{
			"Custom-Header": "hello world",
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	var data map[string]interface{}
	if err := r.JSON(&data); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data["headers"])
}
