package create

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRoundTrip_1(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		t       *transport
		args    args
		wantRes *http.Response
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

			gotRes, err := tt.t.RoundTrip(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoundTrip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("RoundTrip() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
