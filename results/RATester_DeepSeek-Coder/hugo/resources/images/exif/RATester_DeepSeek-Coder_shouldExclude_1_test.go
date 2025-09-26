package exif

import (
	"fmt"
	"testing"
)

func TestShouldExclude_1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		d    *Decoder
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

			if got := tt.d.shouldExclude(tt.args.s); got != tt.want {
				t.Errorf("Decoder.shouldExclude() = %v, want %v", got, tt.want)
			}
		})
	}
}
