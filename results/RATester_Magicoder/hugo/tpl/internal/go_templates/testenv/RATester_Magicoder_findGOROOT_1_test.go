package testenv

import (
	"fmt"
	"testing"
)

func TestfindGOROOT_1(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			want:    "expected GOROOT path",
			wantErr: false,
		},
		{
			name:    "Test case 2",
			want:    "expected GOROOT path",
			wantErr: true,
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

			got, err := findGOROOT()
			if (err != nil) != tt.wantErr {
				t.Errorf("findGOROOT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("findGOROOT() = %v, want %v", got, tt.want)
			}
		})
	}
}
