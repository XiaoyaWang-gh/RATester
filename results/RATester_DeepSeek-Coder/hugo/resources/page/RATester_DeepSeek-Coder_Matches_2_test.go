package page

import (
	"fmt"
	"testing"
)

func TestMatches_2(t *testing.T) {
	type args struct {
		p Page
	}
	tests := []struct {
		name string
		m    PageMatcher
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

			if got := tt.m.Matches(tt.args.p); got != tt.want {
				t.Errorf("Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}
