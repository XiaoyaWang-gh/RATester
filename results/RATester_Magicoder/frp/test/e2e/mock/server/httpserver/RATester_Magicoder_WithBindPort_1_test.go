package httpserver

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithBindPort_1(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			name: "TestWithBindPort",
			args: args{port: 8080},
			want: &Server{bindPort: 8080},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := WithBindPort(tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithBindPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
