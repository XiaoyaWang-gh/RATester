package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMustAsset_1(t *testing.T) {
	tests := []struct {
		name     string
		asset    string
		expected []byte
		wantErr  bool
	}{
		{
			name:     "Test case 1",
			asset:    "test.txt",
			expected: []byte("This is a test file."),
			wantErr:  false,
		},
		{
			name:     "Test case 2",
			asset:    "nonexistent.txt",
			expected: nil,
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

			got, err := Asset(tt.asset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Asset() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Asset() = %v, want %v", got, tt.expected)
			}
		})
	}
}
