package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetStringSlicePreserveString_1(t *testing.T) {
	type args struct {
		cfg Provider
		key string
	}
	tests := []struct {
		name string
		args args
		want []string
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

			if got := GetStringSlicePreserveString(tt.args.cfg, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStringSlicePreserveString() = %v, want %v", got, tt.want)
			}
		})
	}
}
