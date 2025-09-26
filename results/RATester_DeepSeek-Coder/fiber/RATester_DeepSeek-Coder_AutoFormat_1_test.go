package fiber

import (
	"fmt"
	"testing"
)

func TestAutoFormat_1(t *testing.T) {
	t.Parallel()

	app := New()

	tests := []struct {
		name    string
		body    any
		wantErr bool
	}{
		{
			name:    "Test with string body",
			body:    "Hello, World",
			wantErr: false,
		},
		{
			name:    "Test with byte slice body",
			body:    []byte("Hello, World"),
			wantErr: false,
		},
		{
			name:    "Test with int body",
			body:    123,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &DefaultCtx{
				app: app,
			}
			if err := c.AutoFormat(tt.body); (err != nil) != tt.wantErr {
				t.Errorf("AutoFormat() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
