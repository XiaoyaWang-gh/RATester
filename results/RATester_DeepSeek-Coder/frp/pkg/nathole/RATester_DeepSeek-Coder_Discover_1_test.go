package nathole

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestDiscover_1(t *testing.T) {
	type args struct {
		stunServers []string
		localAddr   string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		want1   net.Addr
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

			got, got1, err := Discover(tt.args.stunServers, tt.args.localAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Discover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Discover() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Discover() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
