package fiber

import (
	"fmt"
	"testing"
)

func TestGetPathOriginal_1(t *testing.T) {
	type fields struct {
		pathOriginal string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				pathOriginal: "test_path_original",
			},
			want: "test_path_original",
		},
		{
			name: "Test Case 2",
			fields: fields{
				pathOriginal: "test_path_original_2",
			},
			want: "test_path_original_2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &DefaultCtx{
				pathOriginal: tt.fields.pathOriginal,
			}
			if got := c.getPathOriginal(); got != tt.want {
				t.Errorf("DefaultCtx.getPathOriginal() = %v, want %v", got, tt.want)
			}
		})
	}
}
