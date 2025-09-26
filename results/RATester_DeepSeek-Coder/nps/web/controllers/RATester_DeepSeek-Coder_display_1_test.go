package controllers

import (
	"fmt"
	"testing"
)

func TestDisplay_1(t *testing.T) {
	type args struct {
		tpl []string
	}
	tests := []struct {
		name string
		s    *BaseController
		args args
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

			tt.s.display(tt.args.tpl...)
		})
	}
}
