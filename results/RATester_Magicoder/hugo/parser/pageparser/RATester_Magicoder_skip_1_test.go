package pageparser

import (
	"fmt"
	"testing"
)

func Testskip_1(t *testing.T) {
	type args struct {
		s *sectionHandler
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

			if got := tt.args.s.skip(); got != tt.want {
				t.Errorf("skip() = %v, want %v", got, tt.want)
			}
		})
	}
}
