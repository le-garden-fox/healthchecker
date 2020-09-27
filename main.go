package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// HealthCheckResponse base response
type HealthCheckResponse struct {
	Alive        bool
	Host         string
	Port         string
	ErrorMessage string
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(jsonContentTypeMiddleware)
	headers := handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authrization"})
	methods := handlers.AllowedMethods([]string{"POST", "PUT", "GET", "DELETE", "OPTION"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/", rootLink).Methods("GET")
	router.HandleFunc("/healthcheck-tcp/{host}/{port}", tcpHealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}

func rootLink(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(HealthCheckResponse{
		Alive:        true,
		Host:         "",
		Port:         "",
		ErrorMessage: "",
	})
}

func tcpHealthCheck(w http.ResponseWriter, r *http.Request) {

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
}
