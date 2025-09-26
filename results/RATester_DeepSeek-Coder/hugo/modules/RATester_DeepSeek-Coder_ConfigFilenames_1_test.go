package modules

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfigFilenames_1(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Test case 1",
			want: []string{"config1.yaml", "config2.yaml"},
		},
		{
			name: "Test case 2",
			want: []string{"config3.yaml", "config4.yaml"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &moduleAdapter{
				configFilenames: tt.want,
			}

			got := m.ConfigFilenames()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moduleAdapter.ConfigFilenames() = %v, want %v", got, tt.want)
			}
		})
	}
}
