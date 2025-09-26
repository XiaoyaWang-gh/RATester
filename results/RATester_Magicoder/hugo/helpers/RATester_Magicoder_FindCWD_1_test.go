package helpers

import (
	"fmt"
	"testing"
)

func TestFindCWD_1(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test case 1",
			want:    "/path/to/current/working/directory",
			wantErr: false,
		},
		{
			name:    "Test case 2",
			want:    "/path/to/another/current/working/directory",
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

			got, err := FindCWD()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindCWD() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindCWD() = %v, want %v", got, tt.want)
			}
		})
	}
}
