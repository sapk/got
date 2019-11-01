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
package main

//go:generate go run -mod=vendor github.com/go-swagger/go-swagger/cmd/swagger generate spec -o ./assets/swagger/swagger.v1.json

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog"
)

// PongResponse represent a pong response
// swagger:response PongResponse
type PongResponse struct {
	// The pong response
	// in: body
	Body string
}

func api(webLogger *zerolog.Logger) func(r chi.Router) {
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
			// - application/json
			// responses:
			//   "500":
			//     description: Internal Server Error
			//   "200":
			//     "$ref": "#/responses/PongResponse"

			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode("pong")
			if err != nil {
				webLogger.Warn().Err(err).Msg("Fail to send pong")
			}
		})
	}
}
