package hugo

import (
	"fmt"
	"testing"
)

func TestCompareFloatWithVersion_1(t *testing.T) {
	type args struct {
		v1 float64
		v2 Version
	}
	tests := []struct {
		name string
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

			if got := compareFloatWithVersion(tt.args.v1, tt.args.v2); got != tt.want {
				t.Errorf("compareFloatWithVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
