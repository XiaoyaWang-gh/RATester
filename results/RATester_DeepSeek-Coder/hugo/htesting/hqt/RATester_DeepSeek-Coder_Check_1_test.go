package hqt

import (
	"fmt"
	"testing"
)

func TestCheck_1(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		got     any
		args    []any
		wantErr bool
	}{
		{
			name:    "same strings",
			got:     "test",
			args:    []any{"test"},
			wantErr: false,
		},
		{
			name:    "different strings",
			got:     "test",
			args:    []any{"not test"},
			wantErr: true,
		},
		{
			name:    "same strings with different casing",
			got:     "Test",
			args:    []any{"test"},
			wantErr: false,
		},
		{
			name:    "same strings with different casing and spaces",
			got:     "Test String",
			args:    []any{"test string"},
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

			c := &stringChecker{}
			err := c.Check(tt.got, tt.args, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
