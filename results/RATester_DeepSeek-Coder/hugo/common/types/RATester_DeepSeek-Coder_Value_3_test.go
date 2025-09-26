package types

import (
	"fmt"
	"testing"
)

func TestValue_3(t *testing.T) {
	type args struct {
		source string
		lh     LowHigh[string]
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				source: "Hello, World",
				lh:     LowHigh[string]{Low: 0, High: 5},
			},
			want: "Hello",
		},
		{
			name: "Test case 2",
			args: args{
				source: "Hello, World",
				lh:     LowHigh[string]{Low: 7, High: 11},
			},
			want: "World",
		},
		{
			name: "Test case 3",
			args: args{
				source: "Hello, World",
				lh:     LowHigh[string]{Low: 0, High: 11},
			},
			want: "Hello, World",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.args.lh.Value(tt.args.source); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
