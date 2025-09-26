package hugo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/godartsass/v2"
)

func TestDartSassVersion_1(t *testing.T) {
	tests := []struct {
		name string
		want godartsass.DartSassVersion
	}{
		{
			name: "Test case 1",
			want: godartsass.DartSassVersion{
				ProtocolVersion:       "2.1.0",
				CompilerVersion:       "2.1.0",
				ImplementationVersion: "2.1.0",
				ImplementationName:    "Dart Sass",
				ID:                    1,
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

			if got := dartSassVersion(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dartSassVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
