package allconfig

import (
	"fmt"
	"reflect"
	"testing"
)

func TestconfigFilenames_2(t *testing.T) {
	tests := []struct {
		name     string
		d        ConfigSourceDescriptor
		expected []string
	}{
		{
			name: "empty filename",
			d: ConfigSourceDescriptor{
				Filename: "",
			},
			expected: nil,
		},
		{
			name: "single filename",
			d: ConfigSourceDescriptor{
				Filename: "config.toml",
			},
			expected: []string{"config.toml"},
		},
		{
			name: "multiple filenames",
			d: ConfigSourceDescriptor{
				Filename: "config1.toml,config2.toml,config3.toml",
			},
			expected: []string{"config1.toml", "config2.toml", "config3.toml"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			result := test.d.configFilenames()
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
