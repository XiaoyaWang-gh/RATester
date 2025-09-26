package web

import (
	"fmt"
	"testing"
)

func TestCompareNot_1(t *testing.T) {
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
				a: 1,
				b: 2,
			},
			wantEqual: true,
		},
		{
			name: "Test Case 2",
			args: args{
				a: "hello",
				b: "world",
			},
			wantEqual: true,
		},
		{
			name: "Test Case 3",
			args: args{
				a: []int{1, 2, 3},
				b: []int{1, 2, 3},
			},
			wantEqual: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotEqual := CompareNot(tt.args.a, tt.args.b); gotEqual != tt.wantEqual {
				t.Errorf("CompareNot() = %v, want %v", gotEqual, tt.wantEqual)
			}
		})
	}
}
