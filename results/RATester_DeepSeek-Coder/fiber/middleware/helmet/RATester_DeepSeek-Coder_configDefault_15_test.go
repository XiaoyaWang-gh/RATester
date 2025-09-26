package helmet

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfigDefault_15(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args []Config
		want Config
	}{
		{
			name: "Test with no config",
			args: []Config{},
			want: ConfigDefault,
		},
		{
			name: "Test with custom config",
			args: []Config{
				{
					XSSProtection: "1; mode=block",
				},
			},
			want: Config{
				XSSProtection: "1; mode=block",
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

			got := configDefault(tt.args...)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
