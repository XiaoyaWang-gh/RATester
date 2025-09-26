package static

import (
	"fmt"
	"testing"
)

func TestHasUserDefinedEntrypoint_1(t *testing.T) {
	tests := []struct {
		name string
		c    *Configuration
		want bool
	}{
		{
			name: "has user defined entrypoint",
			c: &Configuration{
				EntryPoints: EntryPoints{
					"http": &EntryPoint{
						Address: ":80",
					},
				},
			},
			want: true,
		},
		{
			name: "no user defined entrypoint",
			c: &Configuration{
				EntryPoints: EntryPoints{},
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

			if got := tt.c.hasUserDefinedEntrypoint(); got != tt.want {
				t.Errorf("Configuration.hasUserDefinedEntrypoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
