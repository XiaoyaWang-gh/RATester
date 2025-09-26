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
		// TODO: Add test cases.
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
