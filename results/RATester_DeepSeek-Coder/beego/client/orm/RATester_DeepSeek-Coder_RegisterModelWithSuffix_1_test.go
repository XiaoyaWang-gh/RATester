package orm

import (
	"fmt"
	"testing"
)

func TestRegisterModelWithSuffix_1(t *testing.T) {
	type testStruct struct {
		name    string
		suffix  string
		models  []interface{}
		wantErr bool
	}

	tests := []testStruct{
		{
			name:    "Test case 1",
			suffix:  "suffix1",
			models:  []interface{}{"model1", "model2"},
			wantErr: false,
		},
		{
			name:    "Test case 2",
			suffix:  "suffix2",
			models:  []interface{}{"model3", "model4"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if tt.wantErr {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("Expected a panic, but no panic occurred")
					}
				}()
			}

			RegisterModelWithSuffix(tt.suffix, tt.models...)
		})
	}
}
