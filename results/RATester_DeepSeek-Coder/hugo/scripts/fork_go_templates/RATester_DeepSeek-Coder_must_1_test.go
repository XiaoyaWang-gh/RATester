package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestMust_1(t *testing.T) {
	type args struct {
		err  error
		what []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 1",
			args: args{
				err:  nil,
				what: []string{"Test"},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				err:  errors.New("Test Error"),
				what: []string{"Test"},
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

			must(tt.args.err, tt.args.what...)
		})
	}
}
