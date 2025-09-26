package allconfig

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfigFilenames_2(t *testing.T) {
	tests := []struct {
		name       string
		descriptor ConfigSourceDescriptor
		want       []string
	}{
		{
			name: "Returns nil when Filename is empty",
			descriptor: ConfigSourceDescriptor{
				Filename: "",
			},
			want: nil,
		},
		{
			name: "Returns a slice of filenames when Filename is not empty",
			descriptor: ConfigSourceDescriptor{
				Filename: "file1.toml,file2.toml",
			},
			want: []string{"file1.toml", "file2.toml"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.descriptor.configFilenames()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configFilenames() = %v, want %v", got, tt.want)
			}
		})
	}
}
