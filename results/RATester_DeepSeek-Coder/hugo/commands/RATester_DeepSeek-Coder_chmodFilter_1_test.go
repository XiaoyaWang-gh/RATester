package commands

import (
	"fmt"
	"os"
	"testing"
)

func TestChmodFilter_1(t *testing.T) {
	type args struct {
		dst os.FileInfo
		src os.FileInfo
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

			if got := chmodFilter(tt.args.dst, tt.args.src); got != tt.want {
				t.Errorf("chmodFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
