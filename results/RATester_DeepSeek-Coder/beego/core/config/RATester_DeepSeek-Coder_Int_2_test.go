package config

import (
	"fmt"
	"testing"
)

func TestInt_2(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		want    int
		wantErr bool
	}{
		{
			name:    "Test case 1",
			key:     "key1",
			want:    1,
			wantErr: false,
		},
		{
			name:    "Test case 2",
			key:     "key2",
			want:    2,
			wantErr: false,
		},
		{
			name:    "Test case 3",
			key:     "key3",
			want:    3,
			wantErr: false,
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

			got, err := Int(tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Int() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
