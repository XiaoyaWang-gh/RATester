package httpcache

import (
	"fmt"
	"testing"
)

func TestIsPollingDisabled_1(t *testing.T) {
	tests := []struct {
		name     string
		compiled *ConfigCompiled
		want     bool
	}{
		{
			name: "all polling disabled",
			compiled: &ConfigCompiled{
				PollConfigs: []PollConfigCompiled{
					{Config: PollConfig{Disable: true}},
					{Config: PollConfig{Disable: true}},
				},
			},
			want: true,
		},
		{
			name: "some polling disabled",
			compiled: &ConfigCompiled{
				PollConfigs: []PollConfigCompiled{
					{Config: PollConfig{Disable: true}},
					{Config: PollConfig{Disable: false}},
				},
			},
			want: false,
		},
		{
			name: "no polling",
			compiled: &ConfigCompiled{
				PollConfigs: []PollConfigCompiled{
					{Config: PollConfig{Disable: false}},
					{Config: PollConfig{Disable: false}},
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

			if got := tt.compiled.IsPollingDisabled(); got != tt.want {
				t.Errorf("IsPollingDisabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
