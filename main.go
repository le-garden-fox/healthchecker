package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/healthcheck-tcp/{host}/{port}", tcpHealthCheck)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
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
