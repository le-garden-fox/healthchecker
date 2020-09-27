package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/le-garden-fox/healthchecker/middleware"
	"github.com/le-garden-fox/healthchecker/routes"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware.JSONContentTypeMiddleware)
	headers := handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authrization"})
	methods := handlers.AllowedMethods([]string{"POST", "PUT", "GET", "DELETE", "OPTION"})
	origins := handlers.AllowedOrigins([]string{"*"})

	routes.RootLink().AddRoute(router)
	routes.HealthCheckTCP().AddRoute(router)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}
