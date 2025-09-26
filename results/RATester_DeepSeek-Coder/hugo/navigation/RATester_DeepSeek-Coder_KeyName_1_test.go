package navigation

import (
	"fmt"
	"testing"
)

func TestKeyName_1(t *testing.T) {
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
		{
			name: "Test KeyName with Identifier",
			fields: fields{
				MenuConfig: MenuConfig{
					Identifier: "TestIdentifier",
				},
			},
			want: "TestIdentifier",
		},
		{
			name: "Test KeyName with Name",
			fields: fields{
				MenuConfig: MenuConfig{
					Name: "TestName",
				},
			},
			want: "TestName",
		},
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
			if got := m.KeyName(); got != tt.want {
				t.Errorf("MenuEntry.KeyName() = %v, want %v", got, tt.want)
			}
		})
	}
}
