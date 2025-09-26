package controllers

import (
	"fmt"
	"testing"
)

func TestGetBoolNoErr_1(t *testing.T) {
	type args struct {
		key string
		def []bool
	}
	tests := []struct {
		name string
		s    *BaseController
		args args
		want bool
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

			if got := tt.s.GetBoolNoErr(tt.args.key, tt.args.def...); got != tt.want {
				t.Errorf("GetBoolNoErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
