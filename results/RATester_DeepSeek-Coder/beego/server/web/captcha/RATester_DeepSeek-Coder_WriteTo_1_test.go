package captcha

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWriteTo_1(t *testing.T) {
	tests := []struct {
		name    string
		img     *Image
		wantN   int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := &bytes.Buffer{}
			gotN, err := tt.img.WriteTo(w)
			if (err != nil) != tt.wantErr {
				t.Errorf("Image.WriteTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Image.WriteTo() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
