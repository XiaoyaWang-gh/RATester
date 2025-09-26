package web

import (
	"fmt"
	"testing"
)

func TestHasTemplateExt_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"Test case 1", "index.tpl", true},
		{"Test case 2", "index.html", false},
		{"Test case 3", "index.tmpl", true},
		{"Test case 4", "index.txt", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasTemplateExt(tt.args); got != tt.want {
				t.Errorf("HasTemplateExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
