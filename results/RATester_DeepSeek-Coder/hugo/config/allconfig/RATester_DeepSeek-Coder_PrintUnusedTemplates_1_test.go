package allconfig

import (
	"fmt"
	"testing"
)

func TestPrintUnusedTemplates_1(t *testing.T) {
	type fields struct {
		config *Config
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test case 1",
			fields: fields{
				config: &Config{
					RootConfig: RootConfig{
						PrintUnusedTemplates: true,
					},
				},
			},
			want: true,
		},
		{
			name: "Test case 2",
			fields: fields{
				config: &Config{
					RootConfig: RootConfig{
						PrintUnusedTemplates: false,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := ConfigLanguage{
				config: tt.fields.config,
			}
			if got := c.PrintUnusedTemplates(); got != tt.want {
				t.Errorf("PrintUnusedTemplates() = %v, want %v", got, tt.want)
			}
		})
	}
}
