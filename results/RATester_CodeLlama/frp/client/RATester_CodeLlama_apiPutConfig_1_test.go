package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiPutConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{
		configFilePath: "./frpc.ini",
	}
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/api/config", nil)
	svr.apiPutConfig(w, r)
	if w.Code != 200 {
		t.Errorf("apiPutConfig failed, code: %d, msg: %s", w.Code, w.Body.String())
	}
}
