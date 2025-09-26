package images

import (
	"fmt"
	"testing"
)

func TestImageFormatFromExt_1(t *testing.T) {
	tests := []struct {
		name      string
		ext       string
		wantF     Format
		wantFound bool
	}{
		{
			name:      "JPEG",
			ext:       ".jpg",
			wantF:     JPEG,
			wantFound: true,
		},
		{
			name:      "PNG",
			ext:       ".png",
			wantF:     PNG,
			wantFound: true,
		},
		{
			name:      "Unknown",
			ext:       ".unknown",
			wantF:     0,
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotF, gotFound := ImageFormatFromExt(tt.ext)
			if gotF != tt.wantF || gotFound != tt.wantFound {
				t.Errorf("ImageFormatFromExt() = %v, %v; want %v, %v", gotF, gotFound, tt.wantF, tt.wantFound)
			}
		})
	}
}
