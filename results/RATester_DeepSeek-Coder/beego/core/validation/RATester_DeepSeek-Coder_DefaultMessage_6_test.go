package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_6(t *testing.T) {
	type fields struct {
		Max int
		Key string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestDefaultMessage",
			fields: fields{
				Max: 10,
				Key: "Max",
			},
			want: "The maximum value is 10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := Max{
				Max: tt.fields.Max,
				Key: tt.fields.Key,
			}
			if got := m.DefaultMessage(); got != tt.want {
				t.Errorf("DefaultMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
