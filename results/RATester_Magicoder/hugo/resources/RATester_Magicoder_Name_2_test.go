package resources

import (
	"fmt"
	"testing"
)

func TestName_2(t *testing.T) {
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
				name: "Test Name",
			},
			want: "Test Name",
		},
		{
			name: "Test Case 2",
			fields: fields{
				name: "Another Test Name",
			},
			want: "Another Test Name",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &genericResource{
				name: tt.fields.name,
			}
			if got := l.Name(); got != tt.want {
				t.Errorf("genericResource.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
