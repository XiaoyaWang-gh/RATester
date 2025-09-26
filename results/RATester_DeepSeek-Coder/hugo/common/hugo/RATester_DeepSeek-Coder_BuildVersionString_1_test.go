package hugo

import (
	"fmt"
	"testing"
)

func TestBuildVersionString_1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test case 1",
			want: "hugo v0.105.0-123abcdef+extended linux/amd64 BuildDate=2022-01-01 VendorInfo=Hugo Inc.",
		},
		{
			name: "Test case 2",
			want: "hugo v0.105.0-123abcdef+extended linux/amd64 BuildDate=2022-01-01",
		},
		{
			name: "Test case 3",
			want: "hugo v0.105.0-123abcdef linux/amd64 BuildDate=2022-01-01",
		},
		{
			name: "Test case 4",
			want: "hugo v0.105.0 linux/amd64 BuildDate=2022-01-01",
		},
		{
			name: "Test case 5",
			want: "hugo v0.105.0 linux/amd64 BuildDate=unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := BuildVersionString(); got != tt.want {
				t.Errorf("BuildVersionString() = %v, want %v", got, tt.want)
			}
		})
	}
}
