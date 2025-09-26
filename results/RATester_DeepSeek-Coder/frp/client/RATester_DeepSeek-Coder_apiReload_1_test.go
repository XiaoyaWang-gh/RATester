package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiReload_1(t *testing.T) {
	t.Run("test apiReload", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		svr := &Service{
			configFilePath: "testdata/frpc.ini",
		}
		req, err := http.NewRequest("GET", "/api/reload", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(svr.apiReload)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})
}
