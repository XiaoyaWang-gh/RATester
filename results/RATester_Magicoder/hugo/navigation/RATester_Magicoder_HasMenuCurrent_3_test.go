package navigation

import (
	"fmt"
	"testing"
)

func TestHasMenuCurrent_3(t *testing.T) {
	type args struct {
		menuID string
		me     *MenuEntry
	}
	tests := []struct {
		name string
		m    nopPageMenus
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

			if got := tt.m.HasMenuCurrent(tt.args.menuID, tt.args.me); got != tt.want {
				t.Errorf("nopPageMenus.HasMenuCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}
