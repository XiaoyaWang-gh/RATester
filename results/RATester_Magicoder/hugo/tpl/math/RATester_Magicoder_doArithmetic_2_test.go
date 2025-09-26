package math

import (
	"fmt"
	"reflect"
	"testing"
)

func TestdoArithmetic_2(t *testing.T) {
	ns := Namespace{}

	tests := []struct {
		name      string
		inputs    []any
		operation rune
		wantValue any
		wantErr   bool
	}{
		{
			name:      "Test with two numbers",
			inputs:    []any{1, 2},
			operation: '+',
			wantValue: 3,
			wantErr:   false,
		},
		{
			name:      "Test with more than two numbers",
			inputs:    []any{1, 2, 3},
			operation: '+',
			wantValue: 6,
			wantErr:   false,
		},
		{
			name:      "Test with one number",
			inputs:    []any{1},
			operation: '+',
			wantValue: nil,
			wantErr:   true,
		},
		{
			name:      "Test with non-numeric values",
			inputs:    []any{"1", "2"},
			operation: '+',
			wantValue: nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			gotValue, err := ns.doArithmetic(tt.inputs, tt.operation)
			if (err != nil) != tt.wantErr {
				t.Errorf("doArithmetic() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("doArithmetic() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
