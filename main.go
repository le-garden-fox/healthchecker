package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	headers := handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type", "Authrization"})
	methods := handlers.AllowedMethods([]string{"POST", "PUT", "GET", "DELETE", "OPTION"})
	origins := handlers.AllowedOrigins([]string{"*"})

	router.HandleFunc("/", rootLink).Methods("GET")
	router.HandleFunc("/healthcheck-tcp/{host}/{port}", tcpHealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))
}

func rootLink(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alive"))
}

func tcpHealthCheck(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	port := params["port"]
	host := params["host"]

	timeout := time.Duration(1 * time.Second)
	_, err := net.DialTimeout("tcp", host+":"+port, timeout)
	if err != nil {
		fmt.Fprintf(w, "%s %s %s\n", host, "not responding", err.Error())
	} else {
		fmt.Fprintf(w, "%s %s %s\n", host, "responding on port:", port)
	}

}
