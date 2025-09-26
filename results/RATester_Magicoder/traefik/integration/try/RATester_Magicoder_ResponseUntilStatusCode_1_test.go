package try

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestResponseUntilStatusCode_1(t *testing.T) {
	type args struct {
		req        *http.Request
		timeout    time.Duration
		statusCode int
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Response
		wantErr bool
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

			got, err := ResponseUntilStatusCode(tt.args.req, tt.args.timeout, tt.args.statusCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResponseUntilStatusCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResponseUntilStatusCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
