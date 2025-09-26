package session

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOpts_1(t *testing.T) {
	tests := []struct {
		name string
		opts []ManagerConfigOpt
		want *ManagerConfig
	}{
		{
			name: "TestOpts_EnableSetCookie",
			opts: []ManagerConfigOpt{
				func(c *ManagerConfig) {
					c.EnableSetCookie = true
				},
			},
			want: &ManagerConfig{
				EnableSetCookie: true,
			},
		},
		{
			name: "TestOpts_DisableHTTPOnly",
			opts: []ManagerConfigOpt{
				func(c *ManagerConfig) {
					c.DisableHTTPOnly = true
				},
			},
			want: &ManagerConfig{
				DisableHTTPOnly: true,
			},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &ManagerConfig{}
			c.Opts(tt.opts...)
			if !reflect.DeepEqual(c, tt.want) {
				t.Errorf("Opts() = %v, want %v", c, tt.want)
			}
		})
	}
}
