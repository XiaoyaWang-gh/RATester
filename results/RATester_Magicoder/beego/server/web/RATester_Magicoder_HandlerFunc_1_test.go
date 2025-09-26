package web

import (
	"fmt"
	"testing"
)

func TestHandlerFunc_1(t *testing.T) {
	ctrl := &Controller{
		methodMapping: map[string]func(){
			"test": func() {},
		},
	}

	tests := []struct {
		name     string
		ctrl     *Controller
		fnname   string
		expected bool
	}{
		{
			name:     "test case 1",
			ctrl:     ctrl,
			fnname:   "test",
			expected: true,
		},
		{
			name:     "test case 2",
			ctrl:     ctrl,
			fnname:   "not_exist",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := tt.ctrl.HandlerFunc(tt.fnname)
			if got != tt.expected {
				t.Errorf("HandlerFunc() = %v, want %v", got, tt.expected)
			}
		})
	}
}
