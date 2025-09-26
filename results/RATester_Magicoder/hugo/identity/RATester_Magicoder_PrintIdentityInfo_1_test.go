package identity

import (
	"fmt"
	"testing"
)

func TestPrintIdentityInfo_1(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				v: "test",
			},
		},
		{
			name: "Test case 2",
			args: args{
				v: 123,
			},
		},
		{
			name: "Test case 3",
			args: args{
				v: struct{}{},
			},
		},
		{
			name: "Test case 4",
			args: args{
				v: []int{1, 2, 3},
			},
		},
		{
			name: "Test case 5",
			args: args{
				v: map[string]int{"one": 1, "two": 2},
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

			PrintIdentityInfo(tt.args.v)
		})
	}
}
