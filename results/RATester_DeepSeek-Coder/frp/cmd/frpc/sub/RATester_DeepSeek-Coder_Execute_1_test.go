package sub

import (
	"fmt"
	"testing"
)

func TestExecute_1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "TestExecute",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Execute()
		})
	}
}
