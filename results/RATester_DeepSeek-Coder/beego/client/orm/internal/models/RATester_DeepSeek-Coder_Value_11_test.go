package models

import (
	"fmt"
	"testing"
)

func TestValue_11(t *testing.T) {
	type fields struct {
		BigIntegerField BigIntegerField
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{
			name: "TestValue_1",
			fields: fields{
				BigIntegerField: -9223372036854775808,
			},
			want: -9223372036854775808,
		},
		{
			name: "TestValue_2",
			fields: fields{
				BigIntegerField: 0,
			},
			want: 0,
		},
		{
			name: "TestValue_3",
			fields: fields{
				BigIntegerField: 9223372036854775807,
			},
			want: 9223372036854775807,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			e := tt.fields.BigIntegerField
			if got := e.Value(); got != tt.want {
				t.Errorf("BigIntegerField.Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
