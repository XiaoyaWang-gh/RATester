package alils

import (
	"fmt"
	"testing"
)

func TestSize_3(t *testing.T) {
	tests := []struct {
		name string
		m    *Log
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.Size(); got != tt.want {
				t.Errorf("Log.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
