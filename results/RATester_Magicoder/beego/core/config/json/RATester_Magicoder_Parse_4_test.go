package json

import (
	"fmt"
	"testing"
)

func TestParse_4(t *testing.T) {
	js := &JSONConfig{}

	tests := []struct {
		name     string
		filename string
		wantErr  bool
	}{
		{
			name:     "valid file",
			filename: "testdata/valid.json",
			wantErr:  false,
		},
		{
			name:     "invalid file",
			filename: "testdata/invalid.json",
			wantErr:  true,
		},
		{
			name:     "non-existent file",
			filename: "testdata/nonexistent.json",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := js.Parse(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
