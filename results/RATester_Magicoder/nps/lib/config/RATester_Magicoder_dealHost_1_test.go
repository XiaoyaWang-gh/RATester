package config

import (
	"fmt"
	"reflect"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestdealHost_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *file.Host
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

			if got := dealHost(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dealHost() = %v, want %v", got, tt.want)
			}
		})
	}
}
