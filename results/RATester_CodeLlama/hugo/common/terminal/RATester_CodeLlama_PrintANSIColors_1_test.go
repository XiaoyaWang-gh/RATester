package terminal

import (
	"fmt"
	"os"
	"testing"
)

func TestPrintANSIColors_1(t *testing.T) {
	type args struct {
		f *os.File
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case 1",
			args: args{
				f: &os.File{},
			},
			want: true,
		},
		{
			name: "case 2",
			args: args{
				f: &os.File{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := PrintANSIColors(tt.args.f); got != tt.want {
				t.Errorf("PrintANSIColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
