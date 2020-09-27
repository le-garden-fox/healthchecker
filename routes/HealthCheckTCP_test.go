package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestHealthCheckTCP(t *testing.T) {
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	HealthCheckTCP().AddRoute(r)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/healthcheck-tcp/api.housecanary.com/80", nil))

	// TODO need to mock out the net.DialTimeout method
	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
	}
}
