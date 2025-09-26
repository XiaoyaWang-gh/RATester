package rules

import (
	"fmt"
	"testing"
)

func TestCheckRule_1(t *testing.T) {
	type testCase struct {
		name    string
		rule    *Tree
		wantErr bool
	}

	tests := []testCase{
		{
			name: "valid rule",
			rule: &Tree{
				Matcher: "Host",
				Value:   []string{"example.com"},
			},
			wantErr: false,
		},
		{
			name: "empty args",
			rule: &Tree{
				Matcher: "Host",
				Value:   []string{},
			},
			wantErr: true,
		},
		{
			name: "no args",
			rule: &Tree{
				Matcher: "Host",
			},
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

			err := CheckRule(tt.rule)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckRule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
