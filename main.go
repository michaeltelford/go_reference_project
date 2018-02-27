package main

import (
	"net/http"

	"github.com/justinas/alice"

	"github.com/michaeltelford/go_reference_project/middleware/logger"
)

func main() {
	chain := alice.New(logger.New).Then(newApp())
	http.ListenAndServe(`:8080`, chain)
}

func newApp() http.Handler {
	app := http.NewServeMux()

	app.HandleFunc(`/healthcheck`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
	app.HandleFunc(`/hello`, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`Hello world`))
	})
	app.Handle(`/`, http.NotFoundHandler())

	return app
}
