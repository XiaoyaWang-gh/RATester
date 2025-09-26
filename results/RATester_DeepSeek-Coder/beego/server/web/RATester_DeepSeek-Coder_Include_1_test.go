package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInclude_1(t *testing.T) {
	type args struct {
		cList []ControllerInterface
	}

	tests := []struct {
		name string
		app  *HttpServer
		args args
		want *HttpServer
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

			if got := tt.app.Include(tt.args.cList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HttpServer.Include() = %v, want %v", got, tt.want)
			}
		})
	}
}
