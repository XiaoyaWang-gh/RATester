package validation

import (
	"fmt"
	"testing"
)

func TestSetDefaultMessage_1(t *testing.T) {
	type args struct {
		msg map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				msg: map[string]string{
					"field name": "value",
				},
			},
		},
		{
			name: "Test case 2",
			args: args{
				msg: map[string]string{},
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

			SetDefaultMessage(tt.args.msg)
		})
	}
}
