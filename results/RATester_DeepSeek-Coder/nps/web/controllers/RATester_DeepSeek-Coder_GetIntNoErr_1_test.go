package controllers

import (
	"fmt"
	"testing"
)

func TestGetIntNoErr_1(t *testing.T) {
	type args struct {
		key string
		def []int
	}
	tests := []struct {
		name string
		s    *BaseController
		args args
		want int
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

			if got := tt.s.GetIntNoErr(tt.args.key, tt.args.def...); got != tt.want {
				t.Errorf("GetIntNoErr() = %v, want %v", got, tt.want)
			}
		})
	}
}
