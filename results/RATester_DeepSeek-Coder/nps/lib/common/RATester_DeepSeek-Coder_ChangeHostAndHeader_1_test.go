package common

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestChangeHostAndHeader_1(t *testing.T) {
	type args struct {
		r         *http.Request
		host      string
		header    string
		addr      string
		addOrigin bool
	}
	tests := []struct {
		name string
		args args
		want *http.Request
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

			ChangeHostAndHeader(tt.args.r, tt.args.host, tt.args.header, tt.args.addr, tt.args.addOrigin)
			if !reflect.DeepEqual(tt.args.r, tt.want) {
				t.Errorf("ChangeHostAndHeader() = %v, want %v", tt.args.r, tt.want)
			}
		})
	}
}
