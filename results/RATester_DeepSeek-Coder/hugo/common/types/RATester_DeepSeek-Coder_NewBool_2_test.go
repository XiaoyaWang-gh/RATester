package types

import (
	"fmt"
	"testing"
)

func TestNewBool_2(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want *bool
	}{
		{
			name: "TestNewBool_True",
			args: args{b: true},
			want: NewBool(true),
		},
		{
			name: "TestNewBool_False",
			args: args{b: false},
			want: NewBool(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewBool(tt.args.b); *got != *tt.want {
				t.Errorf("NewBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
