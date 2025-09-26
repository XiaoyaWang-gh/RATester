package page

import (
	"fmt"
	"testing"
)

func TestWrap_2(t *testing.T) {
	type args struct {
		ss string
	}
	tests := []struct {
		name string
		s    HtmlSummary
		args args
		want string
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

			if got := tt.s.wrap(tt.args.ss); got != tt.want {
				t.Errorf("HtmlSummary.wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}
