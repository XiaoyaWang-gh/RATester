package hugo

import (
	"fmt"
	"testing"
)

func TestCompareVersions_1(t *testing.T) {
	type args struct {
		inVersion Version
		in        any
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

			if got := compareVersions(tt.args.inVersion, tt.args.in); got != tt.want {
				t.Errorf("compareVersions() = %v, want %v", got, tt.want)
			}
		})
	}
}
