package visitor

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestConnectServer_1(t *testing.T) {
	tests := []struct {
		name    string
		v       *visitorHelperImpl
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

			got, err := tt.v.ConnectServer()
			if (err != nil) != tt.wantErr {
				t.Errorf("visitorHelperImpl.ConnectServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("visitorHelperImpl.ConnectServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
