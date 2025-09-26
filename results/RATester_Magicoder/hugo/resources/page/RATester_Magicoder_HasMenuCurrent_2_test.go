package page

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/navigation"
)

func TestHasMenuCurrent_2(t *testing.T) {
	type args struct {
		menuID string
		me     *navigation.MenuEntry
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

			if got := tt.p.HasMenuCurrent(tt.args.menuID, tt.args.me); got != tt.want {
				t.Errorf("HasMenuCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}
