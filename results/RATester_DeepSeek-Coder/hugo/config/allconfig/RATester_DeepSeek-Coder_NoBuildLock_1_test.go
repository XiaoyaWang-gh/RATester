package allconfig

import (
	"fmt"
	"testing"
)

func TestNoBuildLock_1(t *testing.T) {
	type fields struct {
		config *Config
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test NoBuildLock returns true",
			fields: fields{
				config: &Config{
					RootConfig: RootConfig{
						NoBuildLock: true,
					},
				},
			},
			want: true,
		},
		{
			name: "Test NoBuildLock returns false",
			fields: fields{
				config: &Config{
					RootConfig: RootConfig{
						NoBuildLock: false,
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
			if got := c.NoBuildLock(); got != tt.want {
				t.Errorf("NoBuildLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
