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
		want Option
	}{
		{
			name: "TestWithCustomHandler",
			args: args{
				handler: func(conn net.Conn) {
					// do something
				},
			},
			want: func(s *Server) *Server {
				s.handler = func(conn net.Conn) {
					// do something
				}
				return s
			},
		},
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
