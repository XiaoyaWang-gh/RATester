package utils

import (
	"fmt"
	"testing"
)

func TestClear_2(t *testing.T) {
	type args struct {
		f StrTo
	}
	tests := []struct {
		name string
		args args
		want StrTo
	}{
		{
			name: "TestClear",
			args: args{
				f: "test",
			},
			want: StrTo(rune(0x1E)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.args.f.Clear()
			if tt.args.f != tt.want {
				t.Errorf("Clear() = %v, want %v", tt.args.f, tt.want)
			}
		})
	}
}
