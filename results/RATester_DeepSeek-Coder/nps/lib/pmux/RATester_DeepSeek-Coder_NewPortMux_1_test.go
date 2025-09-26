package pmux

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewPortMux_1(t *testing.T) {
	type args struct {
		port        int
		managerHost string
	}
	tests := []struct {
		name string
		args args
		want *PortMux
	}{
		{
			name: "Test case 1",
			args: args{
				port:        8080,
				managerHost: "localhost",
			},
			want: &PortMux{
				port:        8080,
				managerHost: "localhost",
				clientConn:  make(chan *PortConn),
				httpConn:    make(chan *PortConn),
				httpsConn:   make(chan *PortConn),
				managerConn: make(chan *PortConn),
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := NewPortMux(tt.args.port, tt.args.managerHost)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPortMux() = %v, want %v", got, tt.want)
			}
		})
	}
}
