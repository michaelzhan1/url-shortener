package main

import (
	r "github.com/michaelzhan1/url-shortener/internals/runtime"
)

func main() {
	server := r.SetupServer()
	server.ListenAndServe()
}