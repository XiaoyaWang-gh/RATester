package hugo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/godartsass/v2"
)

func TestdartSassVersion_1(t *testing.T) {
	tests := []struct {
		name string
		want godartsass.DartSassVersion
	}{
		{
			name: "Test case 1",
			want: godartsass.DartSassVersion{
				ProtocolVersion:       "1.0",
				CompilerVersion:       "1.0",
				ImplementationVersion: "1.0",
				ImplementationName:    "1.0",
				ID:                    1,
			},
		},
		{
			name: "Test case 2",
			want: godartsass.DartSassVersion{
				ProtocolVersion:       "2.0",
				CompilerVersion:       "2.0",
				ImplementationVersion: "2.0",
				ImplementationName:    "2.0",
				ID:                    2,
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

			if got := dartSassVersion(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dartSassVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
