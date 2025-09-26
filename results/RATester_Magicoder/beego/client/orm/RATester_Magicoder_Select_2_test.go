package orm

import (
	"fmt"
	"testing"
)

func TestSelect_2(t *testing.T) {
	qb := &PostgresQueryBuilder{}

	tests := []struct {
		name   string
		fields []string
		want   string
	}{
		{
			name:   "All fields",
			fields: []string{"*"},
			want:   "SELECT *",
		},
		{
			name:   "Single field",
			fields: []string{"field_name"},
			want:   "SELECT \"field_name\"",
		},
		{
			name:   "Multiple fields",
			fields: []string{"field_name1", "field_name2"},
			want:   "SELECT \"field_name1\",\"field_name2\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			qb.Select(tt.fields...)
			got := qb.String()
			if got != tt.want {
				t.Errorf("Select() = %v, want %v", got, tt.want)
			}
		})
	}
}
