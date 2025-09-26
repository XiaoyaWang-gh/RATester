package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAutoPrefix_2(t *testing.T) {
	type args struct {
		prefix string
		c      ControllerInterface
	}
	tests := []struct {
		name string
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

			if got := AutoPrefix(tt.args.prefix, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AutoPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
