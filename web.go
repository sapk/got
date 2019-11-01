package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"

	"github.com/sapk/got/public/swagger"
	"github.com/sapk/got/public/webapp"
)

func web(debug bool, c *Client, port int) {
	webLogger := setupLogger(debug, "web")

	r := chi.NewRouter()

	//Log setup
	//infoLog := webLogger.Level(zerolog.InfoLevel)
	//r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: false}))
	//r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: &infoLog, NoColor: false}))
	r.Use(RequestLogger(webLogger))

	r.Use(middleware.Recoverer)
	//	r.Use(middleware.ThrottleBacklog(cf.Limits.MaxRequests, cf.Limits.BacklogSize, cf.Limits.BacklogTimeout))
	//  r.Use(middleware.Timeout(cf.Limits.RequestTimeout))
	//	r.Use(middleware.Heartbeat("/ping"))
	//	r.Use(heartbeat.Route("/favicon.ico"))
	/*
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome")) //TODO send ui
		})
	*/

	r.Route("/api", api(webLogger))

	r.Mount("/swagger", http.StripPrefix("/swagger/", http.FileServer(swagger.Swagger)))
	r.Mount("/", http.FileServer(webapp.WebApp))

	webLogger.Info().Msgf("Listening on :%d ...", port)
	//webLogger.Info().Msgf(docgen.JSONRoutesDoc(r))
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

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
