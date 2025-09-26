package port

import (
	"fmt"
	"testing"
)

func TestString_2(t *testing.T) {
	type fields struct {
		name          string
		rangePortFrom int
		rangePortTo   int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				name:          "test",
				rangePortFrom: 0,
				rangePortTo:   0,
			},
			want: "Port_test",
		},
		{
			name: "Test case 2",
			fields: fields{
				name:          "test",
				rangePortFrom: 80,
				rangePortTo:   8080,
			},
			want: "Port_test_Port80_Port8080",
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			builder := &nameBuilder{
				name:          tt.fields.name,
				rangePortFrom: tt.fields.rangePortFrom,
				rangePortTo:   tt.fields.rangePortTo,
			}
			if got := builder.String(); got != tt.want {
				t.Errorf("nameBuilder.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
