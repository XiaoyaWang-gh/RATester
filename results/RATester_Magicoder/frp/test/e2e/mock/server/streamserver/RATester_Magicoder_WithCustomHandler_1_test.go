package streamserver

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestWithCustomHandler_1(t *testing.T) {
	type args struct {
		handler func(net.Conn)
	}
	tests := []struct {
		name string
		args args
		want *Server
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

			if got := WithCustomHandler(tt.args.handler); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCustomHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
