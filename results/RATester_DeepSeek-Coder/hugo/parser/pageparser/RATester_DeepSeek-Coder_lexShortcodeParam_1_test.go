package pageparser

import (
	"fmt"
	"testing"
)

func TestLexShortcodeParam_1(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "Quoted Param",
			input:   `"param"`,
			wantErr: false,
		},
		{
			name:    "Unquoted Param",
			input:   `param`,
			wantErr: false,
		},
		{
			name:    "Quoted Param with equals",
			input:   `"param=value"`,
			wantErr: false,
		},
		{
			name:    "Unquoted Param with equals",
			input:   `param=value`,
			wantErr: false,
		},
		{
			name:    "Quoted Param with escape quote",
			input:   `"param\"value"`,
			wantErr: false,
		},
		{
			name:    "Unquoted Param with escape quote",
			input:   `param\"value`,
			wantErr: true,
		},
		{
			name:    "Quoted Param with escape quote and escaped quote",
			input:   `"\"param\"value"`,
			wantErr: false,
		},
		{
			name:    "Unquoted Param with escape quote and escaped quote",
			input:   `\"param\"value`,
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

			l := &pageLexer{
				input: []byte(tt.input),
			}
			got := lexShortcodeParam(l, false)
			if (got != nil) != tt.wantErr {
				t.Errorf("lexShortcodeParam() error = %v, wantErr %v", got, tt.wantErr)
			}
		})
	}
}
