package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDefaultStrings_7(t *testing.T) {
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

			b := &beegoAppConfig{}
			if got := b.DefaultStrings(tt.args.key, tt.args.defaultVal); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("beegoAppConfig.DefaultStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
