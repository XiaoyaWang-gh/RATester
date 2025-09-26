package streamserver

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_3(t *testing.T) {
	type args struct {
		netType Type
		options []Option
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "test case 1",
			args: args{
				netType: "tcp",
				options: []Option{
					func(s *Server) *Server {
						s.bindAddr = "127.0.0.1"
						return s
					},
				},
			},
			want: &Server{
				netType:  "tcp",
				bindAddr: "127.0.0.1",
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

			if got := New(tt.args.netType, tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
