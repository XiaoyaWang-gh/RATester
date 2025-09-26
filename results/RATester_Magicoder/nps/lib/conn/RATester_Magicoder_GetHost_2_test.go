package conn

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetHost_2(t *testing.T) {
	type args struct {
		s *Conn
	}
	tests := []struct {
		name        string
		args        args
		wantMethod  string
		wantAddress string
		wantRb      []byte
		wantErr     bool
		wantR       *http.Request
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

			gotMethod, gotAddress, gotRb, err, gotR := tt.args.s.GetHost()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMethod != tt.wantMethod {
				t.Errorf("GetHost() gotMethod = %v, want %v", gotMethod, tt.wantMethod)
			}
			if gotAddress != tt.wantAddress {
				t.Errorf("GetHost() gotAddress = %v, want %v", gotAddress, tt.wantAddress)
			}
			if !reflect.DeepEqual(gotRb, tt.wantRb) {
				t.Errorf("GetHost() gotRb = %v, want %v", gotRb, tt.wantRb)
			}
			if !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("GetHost() gotR = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
