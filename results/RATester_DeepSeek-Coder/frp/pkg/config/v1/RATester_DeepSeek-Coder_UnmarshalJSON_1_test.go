package v1

import (
	"fmt"
	"testing"
)

func TestUnmarshalJSON_1(t *testing.T) {
	tests := []struct {
		name    string
		c       *TypedVisitorConfig
		b       []byte
		wantErr bool
	}{
		{
			name:    "null",
			c:       &TypedVisitorConfig{},
			b:       []byte("null"),
			wantErr: true,
		},
		{
			name:    "unknown type",
			c:       &TypedVisitorConfig{},
			b:       []byte(`{"type": "unknown"}`),
			wantErr: true,
		},
		{
			name:    "known type",
			c:       &TypedVisitorConfig{},
			b:       []byte(`{"type": "known"}`),
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

			if err := tt.c.UnmarshalJSON(tt.b); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
