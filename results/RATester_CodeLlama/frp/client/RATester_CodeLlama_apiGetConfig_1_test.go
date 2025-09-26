package client

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestApiGetConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	svr := &Service{}
	svr.configFilePath = "./frpc.ini"
	w := httptest.NewRecorder()
	svr.apiGetConfig(w, nil)
	if w.Code != 200 {
		t.Errorf("Http get response [/api/config], code [%d]", w.Code)
	}
}
