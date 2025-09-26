package captcha

import (
	"fmt"
	"testing"
)

func TestSiphash_1(t *testing.T) {
	type args struct {
		k0 uint64
		k1 uint64
		m  uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
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

			if got := siphash(tt.args.k0, tt.args.k1, tt.args.m); got != tt.want {
				t.Errorf("siphash() = %v, want %v", got, tt.want)
			}
		})
	}
}
