package captcha

import (
	"fmt"
	"image/color"
	"reflect"
	"testing"
)

func TestrandomPalette_1(t *testing.T) {
	tests := []struct {
		name string
		want color.Palette
	}{
		{
			name: "Test case 1",
			want: randomPalette(),
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

			if got := randomPalette(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randomPalette() = %v, want %v", got, tt.want)
			}
		})
	}
}
