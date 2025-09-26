package log

import (
	"fmt"
	"testing"
)

func TestFatal_1(t *testing.T) {
	type args struct {
		v []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				v: []any{1, "test", true},
			},
		},
		{
			name: "Test case 2",
			args: args{
				v: []any{},
			},
		},
		{
			name: "Test case 3",
			args: args{
				v: []any{nil},
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

			Fatal(tt.args.v...)
		})
	}
}
