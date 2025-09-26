package main

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestGetLogWriter_1(t *testing.T) {
	type args struct {
		staticConfiguration *static.Configuration
	}
	tests := []struct {
		name string
		args args
		want io.Writer
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

			if got := getLogWriter(tt.args.staticConfiguration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLogWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
