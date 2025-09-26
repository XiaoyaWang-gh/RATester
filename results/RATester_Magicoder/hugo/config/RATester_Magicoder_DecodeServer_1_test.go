package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDecodeServer_1(t *testing.T) {
	type args struct {
		cfg Provider
	}
	tests := []struct {
		name    string
		args    args
		want    Server
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

			got, err := DecodeServer(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
