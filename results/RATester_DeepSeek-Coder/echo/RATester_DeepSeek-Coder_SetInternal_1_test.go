package echo

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestSetInternal_1(t *testing.T) {
	tests := []struct {
		name    string
		he      *HTTPError
		wantErr error
	}{
		{
			name: "Test with nil error",
			he: &HTTPError{
				Message: "Test error",
				Code:    500,
			},
			wantErr: nil,
		},
		{
			name: "Test with non-nil error",
			he: &HTTPError{
				Message: "Test error",
				Code:    500,
			},
			wantErr: errors.New("internal error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			he := &HTTPError{
				Internal: tt.he.Internal,
				Message:  tt.he.Message,
				Code:     tt.he.Code,
			}
			if got := he.SetInternal(tt.wantErr); !reflect.DeepEqual(got.Internal, tt.wantErr) {
				t.Errorf("HTTPError.SetInternal() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}
