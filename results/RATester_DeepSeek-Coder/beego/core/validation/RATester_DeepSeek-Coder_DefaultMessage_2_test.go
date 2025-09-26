package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_2(t *testing.T) {
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
				Key: "MaxSize",
			},
			want: "The length of %v should be less than or equal to %v",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := MaxSize{
				Max: tt.fields.Max,
				Key: tt.fields.Key,
			}
			if got := m.DefaultMessage(); got != tt.want {
				t.Errorf("MaxSize.DefaultMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
