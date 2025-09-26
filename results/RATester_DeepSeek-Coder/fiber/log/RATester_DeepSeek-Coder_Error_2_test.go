package log

import (
	"fmt"
	"testing"
)

func TestError_2(t *testing.T) {
	t.Run("TestError", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		testCases := []struct {
			name string
			args []any
		}{
			{
				name: "Test case 1",
				args: []any{"test1", 1, true},
			},
			{
				name: "Test case 2",
				args: []any{"test2", 2, false},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Recovered in main", r)
					}
				}()

				Error(tc.args...)
			})
		}
	})
}
