package connection

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestGetBridgeListener_1(t *testing.T) {
	type args struct {
		tp string
	}
	tests := []struct {
		name    string
		args    args
		want    net.Listener
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

			got, err := GetBridgeListener(tt.args.tp)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBridgeListener() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBridgeListener() = %v, want %v", got, tt.want)
			}
		})
	}
}
