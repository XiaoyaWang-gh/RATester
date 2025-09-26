package page

import (
	"fmt"
	"testing"
)

func TestEq_6(t *testing.T) {
	type args struct {
		other any
	}
	tests := []struct {
		name string
		p    *nopPage
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

			if got := tt.p.Eq(tt.args.other); got != tt.want {
				t.Errorf("nopPage.Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}
