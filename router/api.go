// Package main GoT API.
//
// This documentation describes the GoT API.
//
//     Schemes: http, https
//     BasePath: /api
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package router

//go:generate go run -mod=vendor github.com/go-swagger/go-swagger/cmd/swagger generate spec -o ../assets/swagger/swagger.v1.json

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"

	"github.com/sapk/got/modules/mbtiles"
)

// PongResponse represent a pong response
// swagger:response PongResponse
type PongResponse struct {
	// The pong response
	// in: body
	Body string
}

// MapTile represent a png map tile response
// swagger:response MapTile
type MapTile struct {
	// The png map tile response
	// in: body
	Body string //TODO the real content of png file
}

func api(webLogger *zerolog.Logger, c *mbtiles.Client) func(r chi.Router) {
	return func(r chi.Router) {
		//TODO remove only for dev
		// Basic CORS
		cors := cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET"},
			AllowedHeaders: []string{"Accept", "Content-Type", "X-CSRF-Token"},
		})
		r.Use(cors.Handler)

		r.Get("/v1/ping", func(w http.ResponseWriter, r *http.Request) {
			// swagger:operation GET /v1/ping ping
			// ---
			// summary: Echo pong to a ping
			// produces:
			// - text/plain
			// responses:
			//   "500":
			//     description: Internal Server Error
			//   "200":
			//     "$ref": "#/responses/PongResponse"

			w.Header().Set("Content-Type", "text/plain")
			_, err := w.Write([]byte("pong"))
			if err != nil {
				webLogger.Warn().Err(err).Msg("Fail to send pong")
			}
		})

		r.Get("/v1/style.json", func(w http.ResponseWriter, r *http.Request) {
			// swagger:operation GET /v1/style.json style.json
			// ---
			// summary: Return style.json definition
			// produces:
			// - text/plain
			// responses:
			//   "500":
			//     description: Internal Server Error
			//   "200":
			//     "$ref": "#/responses/style.json"

			json, err := c.GetStyle()
			if err != nil {
				webLogger.Warn().Err(err).Msg("Fail to generate map style")
				http.Error(w, err.Error(), 500)
			}

			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write(json)
			if err != nil {
				webLogger.Warn().Err(err).Msg("Fail to send map style")
			}
		})

		r.Get("/v1/tilejson", func(w http.ResponseWriter, r *http.Request) {
			// swagger:operation GET /v1/tilejson tilejson
			// ---
			// summary: Return tilejson definition
			// produces:
			// - text/plain
			// responses:
			//   "500":
			//     description: Internal Server Error
			//   "200":
			//     "$ref": "#/responses/TileJSON"

			json, err := c.GetTileJSON()
			if err != nil {
				webLogger.Warn().Err(err).Msg("Fail to generate TileJSON")
				http.Error(w, err.Error(), 500)
			}

			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write(json)
			if err != nil {
				webLogger.Warn().Err(err).Msg("Fail to send tilejson")
			}
		})

		//TODO cache-control Last-Modified ETag
		r.Get("/v1/tiles/{z:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.{ext}", func(w http.ResponseWriter, r *http.Request) {
			// swagger:operation GET /v1/{z}/{x}/{y}.{ext} getTile
			// ---
			// summary: Map tile
			// produces:
			// - image/png
			// - application/vnd.mapbox-vector-tile
			// - application/x-protobuf
			// parameters:
			// - name: z
			//   in: path
			//   description: the zoom
			//   type: integer
			//   required: true
			// - name: x
			//   in: path
			//   description: the x coord
			//   type: integer
			//   required: true
			// - name: 'y'
			//   in: path
			//   description: the y coord
			//   type: integer
			//   required: true
			// responses:
			//   "500":
			//     description: Internal Server Error
			//   "200":
			//     "$ref": "#/responses/MapTile"

			//TODO return http error 400 bad request
			//TODO retrieve bound of mbtile + return 404 not found/out of bound
			z, err := strconv.Atoi(chi.URLParam(r, "z"))
			if err != nil {
				webLogger.Warn().Err(err).Msg("Invalid z value")
			}
			x, err := strconv.Atoi(chi.URLParam(r, "x"))
			if err != nil {
				webLogger.Warn().Err(err).Msg("Invalid x value")
			}
			y, err := strconv.Atoi(chi.URLParam(r, "y"))
			if err != nil {
				webLogger.Warn().Err(err).Msg("Invalid y value")
			}

			buffer, err := c.GetTile(z, x, y)
			if err != nil {
				webLogger.Warn().Err(err).Msg("Fail to find tile")
			}

			contentType := http.DetectContentType(buffer)
			if contentType == "application/x-gzip" {
				w.Header().Set("Content-Type", "application/x-protobuf")
				w.Header().Set("Content-Encoding", "gzip")
			} else {
				w.Header().Set("Content-Type", contentType)
			}

			_, err = w.Write(buffer)
			if err != nil {
				webLogger.Warn().Err(err).Msg("Fail to send tile")
			}
		})

	}
}
