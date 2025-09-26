package migration

import (
	"fmt"
	"testing"
)

func TestGetCreated_1(t *testing.T) {
	type fields struct {
		Created string
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "TestGetCreated",
			fields: fields{
				Created: "2022-01-01 00:00:00",
			},
			want: 1609459200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := &Migration{
				Created: tt.fields.Created,
			}
			if got := m.GetCreated(); got != tt.want {
				t.Errorf("Migration.GetCreated() = %v, want %v", got, tt.want)
			}
		})
	}
}
