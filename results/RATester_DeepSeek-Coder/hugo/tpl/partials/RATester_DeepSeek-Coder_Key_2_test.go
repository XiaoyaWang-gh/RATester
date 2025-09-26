package partials

import (
	"fmt"
	"testing"
)

func TestKey_2(t *testing.T) {
	tests := []struct {
		name     string
		variants []any
		want     string
	}{
		{
			name:     "test1",
			variants: []any{"variant1", "variant2"},
			want:     "hash_of_test1_variant1_variant2",
		},
		{
			name:     "test2",
			variants: []any{"variant3", "variant4"},
			want:     "hash_of_test2_variant3_variant4",
		},
		{
			name:     "test3",
			variants: nil,
			want:     "test3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			k := partialCacheKey{
				Name:     tt.name,
				Variants: tt.variants,
			}
			if got := k.Key(); got != tt.want {
				t.Errorf("Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
