package testenv

import (
	"fmt"
	"testing"
)

func TesthasSymlink_1(t *testing.T) {
	tests := []struct {
		name       string
		wantOk     bool
		wantReason string
	}{
		{
			name:       "Test case 1",
			wantOk:     true,
			wantReason: "",
		},
		{
			name:       "Test case 2",
			wantOk:     false,
			wantReason: "symlinks unsupported: error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotOk, gotReason := hasSymlink()
			if gotOk != tt.wantOk {
				t.Errorf("hasSymlink() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
			if gotReason != tt.wantReason {
				t.Errorf("hasSymlink() gotReason = %v, want %v", gotReason, tt.wantReason)
			}
		})
	}
}
