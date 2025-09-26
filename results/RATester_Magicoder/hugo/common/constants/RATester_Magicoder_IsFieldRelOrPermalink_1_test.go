package constants

import (
	"fmt"
	"testing"
)

func TestIsFieldRelOrPermalink_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "FieldRelPermalink",
			want: true,
		},
		{
			name: "FieldPermalink",
			want: true,
		},
		{
			name: "RandomField",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsFieldRelOrPermalink(tt.name); got != tt.want {
				t.Errorf("IsFieldRelOrPermalink() = %v, want %v", got, tt.want)
			}
		})
	}
}
