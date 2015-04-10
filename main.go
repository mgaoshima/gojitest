package main

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func hello2(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!, and %s", c.URLParams["name"], c.URLParams["hoge"])
}

func root(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root!!")
}

func main() {
	goji.Get("/", root)
	goji.Get("/hello/:name", hello)
	goji.Get("/hello/:name/:hoge", hello2)
	goji.Serve()
}
