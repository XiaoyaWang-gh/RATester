package log

import (
	"fmt"
	"testing"
)

func TestInfof_1(t *testing.T) {
	tests := []struct {
		name   string
		format string
		v      []interface{}
	}{
		{
			name:   "TestInfof_0",
			format: "This is a test %s",
			v:      []interface{}{"message"},
		},
		{
			name:   "TestInfof_1",
			format: "This is another test %d",
			v:      []interface{}{123},
		},
		{
			name:   "TestInfof_2",
			format: "This is a third test %v",
			v:      []interface{}{"value"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Infof(tt.format, tt.v...)
		})
	}
}
