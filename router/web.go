package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sapk/got/modules/log"
	"github.com/sapk/got/modules/mbtiles"
	"github.com/sapk/got/public/swagger"
	"github.com/sapk/got/public/webapp"
)

//Setup init and start a web router
func Setup(debug bool, c *mbtiles.Client, port int) {
	webLogger := log.NewLogger(debug, "web")

	r := chi.NewRouter()

	//Log setup
	//infoLog := webLogger.Level(zerolog.InfoLevel)
	//r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.New(os.Stdout, "", log.LstdFlags), NoColor: false}))
	//r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: &infoLog, NoColor: false}))
	r.Use(log.RequestLogger(webLogger))

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

	r.Route("/api", api(webLogger, c))

	r.Mount("/swagger", http.StripPrefix("/swagger/", http.FileServer(swagger.Swagger)))
	r.Mount("/", http.FileServer(webapp.WebApp))

	webLogger.Info().Msgf("Listening on :%d ...", port)
	//webLogger.Info().Msgf(docgen.JSONRoutesDoc(r))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), r)
	if err != nil {
		webLogger.Fatal().Err(err).Msgf("Fail to start webserver")
	}
}
