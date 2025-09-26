package resources

import (
	"fmt"
	"testing"
)

func TestTitle_2(t *testing.T) {
	type fields struct {
		title string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				title: "Test Title",
			},
			want: "Test Title",
		},
		{
			name: "Test case 2",
			fields: fields{
				title: "Another Test Title",
			},
			want: "Another Test Title",
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
				title: tt.fields.title,
			}
			if got := l.Title(); got != tt.want {
				t.Errorf("genericResource.Title() = %v, want %v", got, tt.want)
			}
		})
	}
}
