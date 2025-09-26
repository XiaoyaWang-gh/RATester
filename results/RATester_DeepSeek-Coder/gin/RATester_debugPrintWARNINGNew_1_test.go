package gin

import (
	"fmt"
	"testing"
)

func TestDebugPrintWARNINGNew_1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test case 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			debugPrintWARNINGNew()
		})
	}
}
