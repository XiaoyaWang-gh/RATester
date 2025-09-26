package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRegisterPprofHandlers_1(t *testing.T) {
	srv := &Server{
		addr:   "localhost:8080",
		router: mux.NewRouter(),
	}

	srv.registerPprofHandlers()

	testCases := []struct {
		name   string
		path   string
		status int
	}{
		{"cmdline", "/debug/pprof/cmdline", http.StatusOK},
		{"profile", "/debug/pprof/profile", http.StatusOK},
		{"symbol", "/debug/pprof/symbol", http.StatusOK},
		{"trace", "/debug/pprof/trace", http.StatusOK},
		{"index", "/debug/pprof/", http.StatusOK},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			req, err := http.NewRequest(http.MethodGet, "http://"+srv.addr+tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			srv.router.ServeHTTP(rr, req)

			if status := rr.Code; status != tc.status {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tc.status)
			}
		})
	}
}
