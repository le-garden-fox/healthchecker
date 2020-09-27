package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RootLink handles root healthcheck
func RootLink() Handler {
	return Handler{
		Route: func(r *mux.Route) {
			r.Path("/").Methods("GET")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(HealthCheckResponse{
				Alive:        true,
				Host:         "",
				Port:         "",
				ErrorMessage: "",
			})
		},
	}
}
