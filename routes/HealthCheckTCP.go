package routes

import (
	"encoding/json"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// HealthCheckTCP handles checks a tcp domain on given port
func HealthCheckTCP() Handler {

	return Handler{
		Route: func(r *mux.Route) {
			r.Path("/healthcheck-tcp/{host}/{port}").Methods("GET")
		},
		Func: func(w http.ResponseWriter, r *http.Request) {

			params := mux.Vars(r)
			port := params["port"]
			host := params["host"]
			var response HealthCheckResponse

			timeout := time.Duration(1 * time.Second)
			_, err := net.DialTimeout("tcp", host+":"+port, timeout)
			if err != nil {
				response = HealthCheckResponse{
					Alive:        false,
					Host:         host,
					Port:         port,
					ErrorMessage: err.Error(),
				}
			} else {
				response = HealthCheckResponse{
					Alive:        true,
					Host:         host,
					Port:         port,
					ErrorMessage: "",
				}
			}

			json.NewEncoder(w).Encode(response)
		},
	}
}
