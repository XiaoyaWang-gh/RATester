package limiter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconfigDefault_3(t *testing.T) {
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
			name: "Test with custom config",
			args: []Config{
				{
					Max: 10,
				},
			},
			want: Config{
				Max: 10,
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
