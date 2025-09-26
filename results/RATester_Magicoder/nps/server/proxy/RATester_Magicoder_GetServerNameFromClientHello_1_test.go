package proxy

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestGetServerNameFromClientHello_1(t *testing.T) {
	type args struct {
		c net.Conn
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []byte
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

			got, got1 := GetServerNameFromClientHello(tt.args.c)
			if got != tt.want {
				t.Errorf("GetServerNameFromClientHello() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetServerNameFromClientHello() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
