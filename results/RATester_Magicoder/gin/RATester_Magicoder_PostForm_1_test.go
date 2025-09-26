package gin

import (
	"fmt"
	"net/url"
	"testing"
)

func TestPostForm_1(t *testing.T) {
	type test struct {
		name     string
		key      string
		expected string
	}

	tests := []test{
		{"Test1", "key1", "value1"},
		{"Test2", "key2", "value2"},
		{"Test3", "key3", "value3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			ctx := &Context{
				formCache: url.Values{
					tt.key: []string{tt.expected},
				},
			}

			got := ctx.PostForm(tt.key)
			if got != tt.expected {
				t.Errorf("PostForm() = %v, want %v", got, tt.expected)
			}
		})
	}
}
