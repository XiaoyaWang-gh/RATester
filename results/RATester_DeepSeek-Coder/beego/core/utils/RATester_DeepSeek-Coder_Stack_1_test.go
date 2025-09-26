package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestStack_1(t *testing.T) {
	tests := []struct {
		name  string
		skip  int
		want  string
		want1 string
	}{
		{
			name:  "Test case 1",
			skip:  0,
			want:  "at TestStack() [",
			want1: "at runtime.Caller() [",
		},
		{
			name:  "Test case 2",
			skip:  1,
			want:  "at TestStack() [",
			want1: "at runtime.Caller() [",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := Stack(tt.skip, "")
			gotStr := string(got)

			if !strings.Contains(gotStr, tt.want) {
				t.Errorf("Stack() = %v, want %v", gotStr, tt.want)
			}

			if !strings.Contains(gotStr, tt.want1) {
				t.Errorf("Stack() = %v, want %v", gotStr, tt.want1)
			}
		})
	}
}
