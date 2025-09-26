package vhost

import (
	"crypto/tls"
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestReadClientHello_1(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    *tls.ClientHelloInfo
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

			got, err := readClientHello(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("readClientHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readClientHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
