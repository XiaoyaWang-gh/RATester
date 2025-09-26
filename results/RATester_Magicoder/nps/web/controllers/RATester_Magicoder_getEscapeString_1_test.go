package controllers

import (
	"fmt"
	"testing"
)

func TestgetEscapeString_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		s    *BaseController
		args args
		want string
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

			if got := tt.s.getEscapeString(tt.args.key); got != tt.want {
				t.Errorf("getEscapeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
