package main

import "github.com/mahiro72/go-simple-echo-server/presentation/http/router"

func main() {
	r := router.New(3000)
	if err := r.Serve();err != nil {
		panic(err)
	}
}
