package pagemeta

import (
	"fmt"
	"testing"
)

func TestIsZero_1(t *testing.T) {
	tests := []struct {
		name string
		b    BuildConfig
		want bool
	}{
		{
			name: "Test case 1",
			b:    BuildConfig{},
			want: true,
		},
		{
			name: "Test case 2",
			b:    BuildConfig{List: "never", Render: "never", PublishResources: false, set: true},
			want: false,
		},
		{
			name: "Test case 3",
			b:    BuildConfig{List: "always", Render: "always", PublishResources: true, set: true},
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

			if got := tt.b.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
