package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetService_3(t *testing.T) {
	type args struct {
		rw      *httptest.ResponseRecorder
		request *http.Request
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			h := Handler{}
			h.getService(tt.args.rw, tt.args.request)
			if got := tt.args.rw.Result().StatusCode; got != tt.want {
				t.Errorf("getService() = %v, want %v", got, tt.want)
			}
		})
	}
}
