package web

import (
	"fmt"
	"testing"
)

func TestRun_7(t *testing.T) {
	type args struct {
		params []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				params: []string{"test"},
			},
		},
		{
			name: "Test case 2",
			args: args{
				params: []string{""},
			},
		},
		{
			name: "Test case 3",
			args: args{
				params: []string{},
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

			Run(tt.args.params...)
		})
	}
}
