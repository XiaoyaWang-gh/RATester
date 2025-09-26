package internal

import (
	"fmt"
	"testing"
)

func TestValue_1(t *testing.T) {
	type fields struct {
		Name     string
		elements []any
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test with no elements",
			fields: fields{
				Name:     "test",
				elements: []any{},
			},
			want: "test",
		},
		{
			name: "Test with elements",
			fields: fields{
				Name:     "test",
				elements: []any{"element1", "element2"},
			},
			want: "test_element1_element2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			k := ResourceTransformationKey{
				Name:     tt.fields.Name,
				elements: tt.fields.elements,
			}
			if got := k.Value(); got != tt.want {
				t.Errorf("ResourceTransformationKey.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
