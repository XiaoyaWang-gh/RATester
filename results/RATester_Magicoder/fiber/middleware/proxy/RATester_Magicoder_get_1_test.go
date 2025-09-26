package proxy

import (
	"fmt"
	"testing"
)

func Testget_1(t *testing.T) {
	rr := &roundrobin{
		pool: []string{"a", "b", "c"},
	}

	tests := []struct {
		name string
		want string
	}{
		{"test1", "a"},
		{"test2", "b"},
		{"test3", "c"},
		{"test4", "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := rr.get(); got != tt.want {
				t.Errorf("roundrobin.get() = %v, want %v", got, tt.want)
			}
		})
	}
}
