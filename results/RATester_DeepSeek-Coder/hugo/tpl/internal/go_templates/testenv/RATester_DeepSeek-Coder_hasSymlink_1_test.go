package testenv

import (
	"fmt"
	"testing"
)

func TestHasSymlink_1(t *testing.T) {
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
			wantReason: "symlinks unsupported: syscall not supported",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ok, reason := hasSymlink()
			if ok != tt.wantOk {
				t.Errorf("hasSymlink() got = %v, want %v", ok, tt.wantOk)
			}
			if reason != tt.wantReason {
				t.Errorf("hasSymlink() got1 = %v, want %v", reason, tt.wantReason)
			}
		})
	}
}
