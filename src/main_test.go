package main

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("When I call main.newApp", t, func() {
		app := newApp()

		Convey("Then a configured http.Handler is returned", func() {
			So(app, ShouldHaveSameTypeAs, http.DefaultServeMux)
		})
	})
}
