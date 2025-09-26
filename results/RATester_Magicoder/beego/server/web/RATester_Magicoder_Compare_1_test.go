package web

import (
	"fmt"
	"testing"
)

func TestCompare_1(t *testing.T) {
	type args struct {
		a interface{}
		b interface{}
	}
	tests := []struct {
		name      string
		args      args
		wantEqual bool
	}{
		{
			name: "Test Case 1",
			args: args{
				a: "Hello",
				b: "Hello",
			},
			wantEqual: true,
		},
		{
			name: "Test Case 2",
			args: args{
				a: "Hello",
				b: "World",
			},
			wantEqual: false,
		},
		{
			name: "Test Case 3",
			args: args{
				a: 123,
				b: 123,
			},
			wantEqual: true,
		},
		{
			name: "Test Case 4",
			args: args{
				a: 123,
				b: 456,
			},
			wantEqual: false,
		},
		{
			name: "Test Case 5",
			args: args{
				a: nil,
				b: nil,
			},
			wantEqual: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotEqual := Compare(tt.args.a, tt.args.b); gotEqual != tt.wantEqual {
				t.Errorf("Compare() = %v, want %v", gotEqual, tt.wantEqual)
			}
		})
	}
}
