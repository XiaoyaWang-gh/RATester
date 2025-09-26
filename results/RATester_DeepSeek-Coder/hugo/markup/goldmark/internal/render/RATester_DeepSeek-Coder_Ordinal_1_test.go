package render

import (
	"fmt"
	"testing"
)

func TestOrdinal_1(t *testing.T) {
	type fields struct {
		ordinal int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "Test case 1",
			fields: fields{ordinal: 1},
			want:   1,
		},
		{
			name:   "Test case 2",
			fields: fields{ordinal: 2},
			want:   2,
		},
		{
			name:   "Test case 3",
			fields: fields{ordinal: 3},
			want:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &hookBase{
				ordinal: tt.fields.ordinal,
			}
			if got := c.Ordinal(); got != tt.want {
				t.Errorf("hookBase.Ordinal() = %v, want %v", got, tt.want)
			}
		})
	}
}
