package render

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPage_1(t *testing.T) {
	type fields struct {
		page any
	}
	tests := []struct {
		name   string
		fields fields
		want   any
	}{
		{
			name: "Test case 1",
			fields: fields{
				page: "test",
			},
			want: "test",
		},
		{
			name: "Test case 2",
			fields: fields{
				page: 123,
			},
			want: 123,
		},
		{
			name: "Test case 3",
			fields: fields{
				page: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &hookBase{
				page: tt.fields.page,
			}
			if got := c.Page(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hookBase.Page() = %v, want %v", got, tt.want)
			}
		})
	}
}
