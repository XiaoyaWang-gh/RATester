package try

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestDoTryRequest_1(t *testing.T) {
	type args struct {
		request    *http.Request
		timeout    time.Duration
		transport  http.RoundTripper
		conditions []ResponseCondition
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

			got, err := doTryRequest(tt.args.request, tt.args.timeout, tt.args.transport, tt.args.conditions...)
			if (err != nil) != tt.wantErr {
				t.Errorf("doTryRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("doTryRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
