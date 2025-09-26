package allconfig

import (
	"fmt"
	"testing"
)

func TestRunning_1(t *testing.T) {
	type fields struct {
		config *Config
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "TestRunning_RunningTrue",
			fields: fields{
				config: &Config{
					Internal: InternalConfig{
						Running: true,
					},
				},
			},
			want: true,
		},
		{
			name: "TestRunning_RunningFalse",
			fields: fields{
				config: &Config{
					Internal: InternalConfig{
						Running: false,
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
			if got := c.Running(); got != tt.want {
				t.Errorf("Running() = %v, want %v", got, tt.want)
			}
		})
	}
}
