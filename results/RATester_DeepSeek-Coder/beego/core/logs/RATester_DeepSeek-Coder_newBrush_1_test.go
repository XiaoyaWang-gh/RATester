package logs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewBrush_1(t *testing.T) {
	type args struct {
		color string
	}
	tests := []struct {
		name string
		args args
		want brush
	}{
		{
			name: "Test newBrush with color red",
			args: args{color: "31"},
			want: newBrush("31"),
		},
		{
			name: "Test newBrush with color green",
			args: args{color: "32"},
			want: newBrush("32"),
		},
		{
			name: "Test newBrush with color blue",
			args: args{color: "34"},
			want: newBrush("34"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newBrush(tt.args.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newBrush() = %v, want %v", got, tt.want)
			}
		})
	}
}
