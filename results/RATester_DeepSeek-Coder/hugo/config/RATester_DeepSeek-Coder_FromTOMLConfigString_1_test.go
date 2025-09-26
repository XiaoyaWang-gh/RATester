package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFromTOMLConfigString_1(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name string
		args args
		want Provider
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

			if got := FromTOMLConfigString(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromTOMLConfigString() = %v, want %v", got, tt.want)
			}
		})
	}
}
