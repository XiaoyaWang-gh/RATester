package navigation

import (
	"fmt"
	"testing"
)

func TestisSameAsDescendantMenu_1(t *testing.T) {
	type args struct {
		inme   *MenuEntry
		parent *MenuEntry
	}
	tests := []struct {
		name string
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

			pm := &pageMenus{}
			if got := pm.isSameAsDescendantMenu(tt.args.inme, tt.args.parent); got != tt.want {
				t.Errorf("pageMenus.isSameAsDescendantMenu() = %v, want %v", got, tt.want)
			}
		})
	}
}
