package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDealCommon_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *CommonConfig
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

			if got := dealCommon(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dealCommon() = %v, want %v", got, tt.want)
			}
		})
	}
}
