package validation

import (
	"fmt"
	"testing"
)

func TestError_2(t *testing.T) {
	type fields struct {
		Message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				Message: "Test message",
			},
			want: "Test message",
		},
		{
			name: "Test case 2",
			fields: fields{
				Message: "",
			},
			want: "",
		},
		{
			name: "Test case 3",
			fields: fields{
				Message: "Another test message",
			},
			want: "Another test message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			e := &Error{
				Message: tt.fields.Message,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
