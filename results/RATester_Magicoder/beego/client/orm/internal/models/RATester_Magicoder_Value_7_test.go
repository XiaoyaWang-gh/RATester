package models

import (
	"fmt"
	"testing"
)

func TestValue_7(t *testing.T) {
	type fields struct {
		value int16
	}
	tests := []struct {
		name   string
		fields fields
		want   int16
	}{
		{
			name: "Test case 1",
			fields: fields{
				value: 10,
			},
			want: 10,
		},
		{
			name: "Test case 2",
			fields: fields{
				value: -10,
			},
			want: -10,
		},
		{
			name: "Test case 3",
			fields: fields{
				value: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			e := SmallIntegerField(tt.fields.value)
			if got := e.Value(); got != tt.want {
				t.Errorf("SmallIntegerField.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
