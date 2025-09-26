package fiber

import (
	"fmt"
	"testing"
)

func TestGetTreePath_1(t *testing.T) {
	type fields struct {
		treePath string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Case 1",
			fields: fields{
				treePath: "test_tree_path",
			},
			want: "test_tree_path",
		},
		{
			name: "Test Case 2",
			fields: fields{
				treePath: "another_tree_path",
			},
			want: "another_tree_path",
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
				treePath: tt.fields.treePath,
			}
			if got := c.getTreePath(); got != tt.want {
				t.Errorf("DefaultCtx.getTreePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
