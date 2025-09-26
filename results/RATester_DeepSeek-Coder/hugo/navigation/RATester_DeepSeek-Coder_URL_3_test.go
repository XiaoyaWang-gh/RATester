package navigation

import (
	"fmt"
	"testing"
)

func TestURL_3(t *testing.T) {
	type fields struct {
		MenuConfig    MenuConfig
		Menu          string
		ConfiguredURL string
		Page          Page
		Children      Menu
	}
	tests := []struct {
		name   string
		fields fields
		want   string
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

			m := &MenuEntry{
				MenuConfig:    tt.fields.MenuConfig,
				Menu:          tt.fields.Menu,
				ConfiguredURL: tt.fields.ConfiguredURL,
				Page:          tt.fields.Page,
				Children:      tt.fields.Children,
			}
			if got := m.URL(); got != tt.want {
				t.Errorf("MenuEntry.URL() = %v, want %v", got, tt.want)
			}
		})
	}
}
