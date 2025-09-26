package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/navigation"
)

func TestIsMenuCurrent_1(t *testing.T) {
	type args struct {
		menuID string
		inme   *navigation.MenuEntry
	}
	tests := []struct {
		name string
		p    *pageMenus
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

			if got := tt.p.IsMenuCurrent(tt.args.menuID, tt.args.inme); got != tt.want {
				t.Errorf("pageMenus.IsMenuCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}
