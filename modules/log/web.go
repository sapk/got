package log

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

// RequestLogger returns a logger handler using a custom LogFormatter.
func RequestLogger(log *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			t1 := time.Now()
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			scheme := "http"
			if r.TLS != nil {
				scheme = "https"
			}
			//TODO color ?
			//TODO better perf
			defer func() {
				log.Info().Msgf(`"GET %s://%s%s %s" from %s - %d %dB in %s`, scheme, r.Host, r.RequestURI, r.Proto, r.RemoteAddr, ww.Status(), ww.BytesWritten(), time.Since(t1))
			}()

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
