package page

import (
	"fmt"
	"testing"
)

func TestsearchPage_1(t *testing.T) {
	type args struct {
		p     Page
		pages Pages
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

			if got := searchPage(tt.args.p, tt.args.pages); got != tt.want {
				t.Errorf("searchPage() = %v, want %v", got, tt.want)
			}
		})
	}
}
