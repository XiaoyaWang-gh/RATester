package constants

import (
	"fmt"
	"testing"
)

func TestIsResourceTransformationPermalinkHash_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{
			name: "ResourceTransformationFingerprint",
			want: true,
		},
		{
			name: "SomeOtherString",
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

			if got := IsResourceTransformationPermalinkHash(tt.name); got != tt.want {
				t.Errorf("IsResourceTransformationPermalinkHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
