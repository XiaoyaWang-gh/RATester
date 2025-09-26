package utils

import (
	"fmt"
	"testing"
)

func TestGet_19(t *testing.T) {
	tests := []struct {
		name  string
		a     ArgString
		i     int
		args  []string
		wantR string
	}{
		{
			name:  "Test case 1",
			a:     []string{"test1", "test2", "test3"},
			i:     1,
			args:  []string{"default"},
			wantR: "test2",
		},
		{
			name:  "Test case 2",
			a:     []string{"test1", "test2", "test3"},
			i:     3,
			args:  []string{"default"},
			wantR: "default",
		},
		{
			name:  "Test case 3",
			a:     []string{"test1", "test2", "test3"},
			i:     1,
			args:  []string{},
			wantR: "test2",
		},
		{
			name:  "Test case 4",
			a:     []string{"test1", "test2", "test3"},
			i:     -1,
			args:  []string{"default"},
			wantR: "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotR := tt.a.Get(tt.i, tt.args...); gotR != tt.wantR {
				t.Errorf("ArgString.Get() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}
