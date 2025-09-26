package accesslog

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInitServiceFields_1(t *testing.T) {
	type args struct {
		rw   *httptest.ResponseRecorder
		req  *http.Request
		next http.Handler
		data *LogData
	}

	tests := []struct {
		name string
		args args
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

			InitServiceFields(tt.args.rw, tt.args.req, tt.args.next, tt.args.data)
		})
	}
}
