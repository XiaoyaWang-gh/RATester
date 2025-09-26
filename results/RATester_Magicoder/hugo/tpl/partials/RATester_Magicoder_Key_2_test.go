package partials

import (
	"fmt"
	"testing"
)

func TestKey_2(t *testing.T) {
	type fields struct {
		Name     string
		Variants []any
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				Name:     "test",
				Variants: []any{1, "2", 3.0},
			},
			want: "test123",
		},
		{
			name: "Test case 2",
			fields: fields{
				Name:     "test2",
				Variants: []any{},
			},
			want: "test2",
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
				Name:     tt.fields.Name,
				Variants: tt.fields.Variants,
			}
			if got := k.Key(); got != tt.want {
				t.Errorf("Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
