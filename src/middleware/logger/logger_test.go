package logger

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	// Only pass t into top-level Convey calls
	Convey("Given I have a mock HTTP app", t, func() {
		app := http.NewServeMux()

		Convey("When I call logger.New", func() {
			middleware := New(app)

			Convey("Then a configured logger middleware is returned", func() {
				So(middleware, ShouldHaveSameTypeAs, logger{})
			})
		})
	})
}

func TestServeHTTP(t *testing.T) {
	Convey("Given I have a mock HTTP app", t, func() {
		app := http.NewServeMux()
		app.HandleFunc(`/hello`, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte(`Hello world`))
		})

		w := httptest.NewRecorder()
		r, _ := http.NewRequest("GET", "/hello", nil)

		var buf bytes.Buffer
		expectedLogOutput := `2018/03/01 : /hello - Request received
2018/03/01 : /hello - Response sent
`

		Convey("And I have a mock logger middleware", func() {
			middleware := logger{
				Next: app,
				Log:  log.New(&buf, ``, 1),
			}

			Convey("When I call ServeHTTP", func() {
				middleware.ServeHTTP(w, r)

				Convey("Then a response status of 200 is returned", func() {
					So(w.Code, ShouldEqual, 200)
				})
				Convey("And a response body of `Hello world` is returned", func() {
					So(w.Body.String(), ShouldEqual, `Hello world`)
				})
				Convey("And the request and response has been logged", func() {
					So(string(buf.Bytes()), ShouldEqual, expectedLogOutput)
				})
			})
		})
	})
}
