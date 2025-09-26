package main

import (
	"errors"
	"fmt"
	"testing"
)

func Testmust_5(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				err: nil,
			},
		},
		{
			name: "Test case 2",
			args: args{
				err: errors.New("some error"),
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

			must(tt.args.err)
		})
	}
}
