package group

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestCreateConnByEndpoint_1(t *testing.T) {
	type args struct {
		endpoint   string
		remoteAddr string
	}
	tests := []struct {
		name    string
		g       *HTTPGroup
		args    args
		want    net.Conn
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

			got, err := tt.g.createConnByEndpoint(tt.args.endpoint, tt.args.remoteAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("HTTPGroup.createConnByEndpoint() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPGroup.createConnByEndpoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
