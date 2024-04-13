package helpers

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"rest/handlers"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// marinas
	router.HandleFunc("/marinas", handlers.ListMarinas).Methods("GET")
	router.HandleFunc("/marinas", handlers.CreateMarina).Methods("POST")
	router.HandleFunc("/marinas/{id}", handlers.GetMarina).Methods("GET")
	router.HandleFunc("/marinas/{id}", handlers.UpdateMarina).Methods("PUT")
	router.HandleFunc("/marinas/{id}", handlers.DeleteMarina).Methods("DELETE")
	router.HandleFunc("/marinas/{id}/yachts", handlers.ListYachtsInMarina).Methods("GET")

	// yachts
	router.HandleFunc("/yachts", handlers.ListYachts).Methods("GET")
	router.HandleFunc("/yachts", handlers.CreateYacht).Methods("POST")
	router.HandleFunc("/yachts/{id}", handlers.GetYacht).Methods("GET")
	router.HandleFunc("/yachts/{id}", handlers.UpdateYacht).Methods("PUT")
	router.HandleFunc("/yachts/{id}", handlers.DeleteYacht).Methods("DELETE")

	// charters
	router.HandleFunc("/charters", handlers.ListCharters).Methods("GET")
	router.HandleFunc("/charters", handlers.CreateCharter).Methods("POST")
	router.HandleFunc("/charters/{id}", handlers.GetCharter).Methods("GET")
	router.HandleFunc("/charters/{id}", handlers.UpdateCharter).Methods("PUT")
	router.HandleFunc("/charters/{id}", handlers.DeleteCharter).Methods("DELETE")

	// migrations
	router.HandleFunc("/migrations", handlers.ListMigrations).Methods("GET")
	router.HandleFunc("/migrations", handlers.CreateMigration).Methods("POST")
	router.HandleFunc("/migrations/{id}", handlers.GetMigration).Methods("GET")

	// tokens
	router.HandleFunc("/tokens", handlers.CreateToken).Methods("POST")

	return router
}

func Request(t *testing.T, method string, url string, body string) *httptest.ResponseRecorder {
	return RequestWithHeaders(t, method, url, body, nil)
}

func RequestWithHeaders(t *testing.T, method string, url string, body string, headers map[string]string) *httptest.ResponseRecorder {
	var req *http.Request
	var err error

	if body != "" {
		req, err = http.NewRequest(method, url, strings.NewReader(body))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	rr := httptest.NewRecorder()
	router := SetupRouter()
	router.ServeHTTP(rr, req)

	return rr
}

func GetRequest(t *testing.T, url string) *httptest.ResponseRecorder {
	rr := Request(t, "GET", url, "")

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("should return 200, got %v", status)
	}

	return rr
}

func CreateRequest(t *testing.T, url string, body string) *httptest.ResponseRecorder {
	return Request(t, "POST", url, body)
}

func DeleteRequest(t *testing.T, url string) *httptest.ResponseRecorder {
	rr := Request(t, "DELETE", url, "")

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("should return 204, got %v", status)
	}

	return rr
}

func UpdateRequest(t *testing.T, url string, body string) *httptest.ResponseRecorder {
	return Request(t, "PUT", url, body)
}

func UpdateRequestWithHeaders(t *testing.T, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := SetupRouter()
	router.ServeHTTP(rr, req)

	return rr
}
