package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefaultStrings_3(t *testing.T) {
	type args struct {
		key        string
		defaultVal []string
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

			if got := DefaultStrings(tt.args.key, tt.args.defaultVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
