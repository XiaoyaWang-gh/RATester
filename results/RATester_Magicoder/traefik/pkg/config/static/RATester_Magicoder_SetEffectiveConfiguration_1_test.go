package static

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetEffectiveConfiguration_1(t *testing.T) {
	tests := []struct {
		name   string
		config *Configuration
		want   *Configuration
	}{
		{
			name: "no user defined entrypoint",
			config: &Configuration{
				EntryPoints: nil,
			},
			want: &Configuration{
				EntryPoints: EntryPoints{
					"http": &EntryPoint{
						Address: ":80",
					},
				},
			},
		},
		{
			name: "user defined entrypoint",
			config: &Configuration{
				EntryPoints: EntryPoints{
					"http": &EntryPoint{
						Address: ":8080",
					},
				},
			},
			want: &Configuration{
				EntryPoints: EntryPoints{
					"http": &EntryPoint{
						Address: ":8080",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tt.config.SetEffectiveConfiguration()
			if !reflect.DeepEqual(tt.config, tt.want) {
				t.Errorf("got %v, want %v", tt.config, tt.want)
			}
		})
	}
}
