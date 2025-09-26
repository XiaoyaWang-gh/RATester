package port

import (
	"fmt"
	"testing"
)

func TestGetByRange_1(t *testing.T) {
	type args struct {
		from int
		to   int
	}
	tests := []struct {
		name string
		pa   *Allocator
		args args
		want int
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

			if got := tt.pa.getByRange(tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("Allocator.getByRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
