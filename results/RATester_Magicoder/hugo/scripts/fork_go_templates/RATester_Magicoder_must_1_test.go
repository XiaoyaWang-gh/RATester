package main

import (
	"errors"
	"fmt"
	"testing"
)

func Testmust_1(t *testing.T) {
	type args struct {
		err  error
		what []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				err:  nil,
				what: []string{"Test"},
			},
		},
		{
			name: "Test case 2",
			args: args{
				err:  errors.New("test error"),
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
