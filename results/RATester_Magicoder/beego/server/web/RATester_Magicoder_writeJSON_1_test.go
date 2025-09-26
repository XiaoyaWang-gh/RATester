package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestwriteJSON_1(t *testing.T) {
	type args struct {
		rw       http.ResponseWriter
		jsonData []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test writeJSON",
			args: args{
				rw:       httptest.NewRecorder(),
				jsonData: []byte(`{"name": "John Doe"}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			writeJSON(tt.args.rw, tt.args.jsonData)
			resp := tt.args.rw.(*httptest.ResponseRecorder).Result()
			body, _ := ioutil.ReadAll(resp.Body)
			var data map[string]interface{}
			json.Unmarshal(body, &data)
			if data["name"] != "John Doe" {
				t.Errorf("Expected name to be 'John Doe', but got %s", data["name"])
			}
		})
	}
}
