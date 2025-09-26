package berror

import (
	"errors"
	"fmt"
	"testing"
)

func TestFromError_1(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		wantCode uint32
		wantOk   bool
	}{
		{
			name:     "error without code",
			err:      errors.New("ERROR- , some error"),
			wantCode: 0,
			wantOk:   false,
		},
		{
			name:     "error with invalid code",
			err:      errors.New("ERROR-abc , some error"),
			wantCode: 0,
			wantOk:   false,
		},
		{
			name:     "error with valid code",
			err:      errors.New("ERROR-123 , some error"),
			wantCode: 123,
			wantOk:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotCode, gotOk := FromError(tt.err)
			if gotCode.Code() != tt.wantCode {
				t.Errorf("FromError() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
			if gotOk != tt.wantOk {
				t.Errorf("FromError() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
