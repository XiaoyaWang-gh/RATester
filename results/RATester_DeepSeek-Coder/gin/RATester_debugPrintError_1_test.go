package gin

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestDebugPrintError_1(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "Test case 1",
			err:  errors.New("test error"),
			want: "[GIN-debug] [ERROR] test error\n",
		},
		{
			name: "Test case 2",
			err:  nil,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			buf := &bytes.Buffer{}
			DefaultErrorWriter = buf
			debugPrintError(tt.err)
			got := buf.String()
			if got != tt.want {
				t.Errorf("debugPrintError() = %v, want %v", got, tt.want)
			}
		})
	}
}
