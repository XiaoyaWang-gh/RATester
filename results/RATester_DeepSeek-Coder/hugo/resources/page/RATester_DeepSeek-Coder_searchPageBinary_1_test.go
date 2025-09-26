package page

import (
	"fmt"
	"testing"
)

func TestSearchPageBinary_1(t *testing.T) {
	type args struct {
		p     Page
		pages Pages
		less  func(p1, p2 Page) bool
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

			if got := searchPageBinary(tt.args.p, tt.args.pages, tt.args.less); got != tt.want {
				t.Errorf("searchPageBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
