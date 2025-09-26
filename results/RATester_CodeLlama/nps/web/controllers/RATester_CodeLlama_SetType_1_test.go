package controllers

import (
	"fmt"
	"testing"
)

func TestSetType_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				name: "test1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &BaseController{}
			s.SetType(tt.args.name)
		})
	}
}
