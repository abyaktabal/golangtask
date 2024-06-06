package main

import (
	"net/http"

	"server.go/router"
)

func main() {
	r := router.NewRouter()
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
