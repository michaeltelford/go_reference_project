package logger

import (
	"log"
	"net/http"
	"os"
)

type logger struct {
	Next http.Handler
	Log  *log.Logger
}

func (l logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.Log.Printf("%s: %s - Request received", r.RemoteAddr, r.URL)
	l.Next.ServeHTTP(w, r)
	l.Log.Printf("%s: %s - Response sent", r.RemoteAddr, r.URL)
}

// New is a constructor function for logging middleware
func New(h http.Handler) http.Handler {
	return logger{
		Next: h,
		Log:  log.New(os.Stdout, ``, 1),
	}
}
