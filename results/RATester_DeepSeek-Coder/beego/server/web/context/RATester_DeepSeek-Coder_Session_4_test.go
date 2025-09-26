package context

import (
	"fmt"
	"testing"
)

func TestSession_4(t *testing.T) {
	type args struct {
		name  interface{}
		value interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestSession_0",
			args: args{
				name:  "test_name",
				value: "test_value",
			},
		},
		{
			name: "TestSession_1",
			args: args{
				name:  123,
				value: 456,
			},
		},
		{
			name: "TestSession_2",
			args: args{
				name:  "test_name",
				value: nil,
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

			output := &BeegoOutput{}
			output.Session(tt.args.name, tt.args.value)
			val := output.Context.Input.CruSession.Get(nil, tt.args.name)
			if val != tt.args.value {
				t.Errorf("Session() = %v, want %v", val, tt.args.value)
			}
		})
	}
}
