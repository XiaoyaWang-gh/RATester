package template

import (
	"fmt"
	"testing"
)

func TestName_7(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				name: "TestName",
			},
			want: "TestName",
		},
		{
			name: "Test Case 2",
			fields: fields{
				name: "TestName2",
			},
			want: "TestName2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tmpl := &Template{
				name: tt.fields.name,
			}
			if got := tmpl.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
