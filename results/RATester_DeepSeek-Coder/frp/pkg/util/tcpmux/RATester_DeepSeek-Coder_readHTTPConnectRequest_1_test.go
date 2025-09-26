package tcpmux

import (
	"fmt"
	"io"
	"testing"
)

func TestReadHTTPConnectRequest_1(t *testing.T) {
	type args struct {
		rd io.Reader
	}
	tests := []struct {
		name         string
		args         args
		wantHost     string
		wantHttpUser string
		wantHttpPwd  string
		wantErr      bool
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

			muxer := &HTTPConnectTCPMuxer{}
			gotHost, gotHttpUser, gotHttpPwd, err := muxer.readHTTPConnectRequest(tt.args.rd)
			if (err != nil) != tt.wantErr {
				t.Errorf("readHTTPConnectRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHost != tt.wantHost {
				t.Errorf("readHTTPConnectRequest() gotHost = %v, want %v", gotHost, tt.wantHost)
			}
			if gotHttpUser != tt.wantHttpUser {
				t.Errorf("readHTTPConnectRequest() gotHttpUser = %v, want %v", gotHttpUser, tt.wantHttpUser)
			}
			if gotHttpPwd != tt.wantHttpPwd {
				t.Errorf("readHTTPConnectRequest() gotHttpPwd = %v, want %v", gotHttpPwd, tt.wantHttpPwd)
			}
		})
	}
}
