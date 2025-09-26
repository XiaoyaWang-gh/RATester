package helmet

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconfigDefault_15(t *testing.T) {
	tests := []struct {
		name string
		args []Config
		want Config
	}{
		{
			name: "Test with no args",
			args: []Config{},
			want: ConfigDefault,
		},
		{
			name: "Test with args",
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

			if got := configDefault(tt.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
