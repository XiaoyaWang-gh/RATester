package logs

import (
	"fmt"
	"testing"
)

func TestSetLogFuncCall_1(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{b: true},
		},
		{
			name: "Test case 2",
			args: args{b: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			SetLogFuncCall(tt.args.b)
			// Add assertions here
		})
	}
}
