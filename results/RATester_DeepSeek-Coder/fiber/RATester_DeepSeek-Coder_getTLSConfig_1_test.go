package fiber

import (
	"crypto/tls"
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestGetTLSConfig_1(t *testing.T) {
	type args struct {
		ln net.Listener
	}
	tests := []struct {
		name string
		args args
		want *tls.Config
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

			if got := getTLSConfig(tt.args.ln); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTLSConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
