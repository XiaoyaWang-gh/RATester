package commands

import (
	"fmt"
	"testing"
)

func TestprintChangeDetected_1(t *testing.T) {
	type args struct {
		typ string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				typ: "test",
			},
		},
		{
			name: "Test case 2",
			args: args{
				typ: "",
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

			c := &hugoBuilder{}
			c.printChangeDetected(tt.args.typ)
		})
	}
}
