package main

import (
	"fmt"
	"log"

	"github.com/tbdsux/reqgo"
)

func main() {
	r, err := reqgo.Post("https://httpbin.org/post", &reqgo.Options{
		Data: map[string]string{
			"hello": "world",
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	var data map[string]interface{}
	if err := r.JSON(&data); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}
