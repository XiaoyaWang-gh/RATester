package hugo

import (
	"fmt"
	"testing"
)

func TestCompare_1(t *testing.T) {
	type args struct {
		other any
	}
	tests := []struct {
		name string
		h    VersionString
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

			if got := tt.h.Compare(tt.args.other); got != tt.want {
				t.Errorf("VersionString.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
