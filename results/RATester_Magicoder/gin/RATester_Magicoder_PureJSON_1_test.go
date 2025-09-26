package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestPureJSON_1(t *testing.T) {
	tests := []struct {
		name string
		code int
		obj  any
	}{
		{
			name: "Test case 1",
			code: http.StatusOK,
			obj:  map[string]string{"key": "value"},
		},
		{
			name: "Test case 2",
			code: http.StatusBadRequest,
			obj:  "error message",
		},
		{
			name: "Test case 3",
			code: http.StatusInternalServerError,
			obj:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Context{}
			c.PureJSON(tt.code, tt.obj)
			// Add assertions here to check the expected behavior
		})
	}
}
