package web

import (
	"fmt"
	"testing"
)

func TestHandlerFunc_1(t *testing.T) {
	ctrl := &Controller{
		methodMapping: map[string]func(){
			"test": func() {
				fmt.Println("test function")
			},
		},
	}

	tests := []struct {
		name   string
		c      *Controller
		fnname string
		want   bool
	}{
		{
			name:   "TestHandlerFunc_0",
			c:      ctrl,
			fnname: "test",
			want:   true,
		},
		{
			name:   "TestHandlerFunc_1",
			c:      ctrl,
			fnname: "not_exist",
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.HandlerFunc(tt.fnname); got != tt.want {
				t.Errorf("HandlerFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
