package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestdebugPrintError_1(t *testing.T) {
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
				err: errors.New("test error"),
			},
		},
		{
			name: "Test case 2",
			args: args{
				err: nil,
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

			debugPrintError(tt.args.err)
		})
	}
}
