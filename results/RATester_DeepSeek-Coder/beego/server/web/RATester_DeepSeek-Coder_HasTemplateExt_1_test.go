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
		{"Test1", "test.tmpl", true},
		{"Test2", "test.html", false},
		{"Test3", "test.tpl", true},
		{"Test4", "test.txt", false},
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
