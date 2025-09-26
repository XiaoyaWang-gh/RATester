package rules

import (
	"fmt"
	"testing"
)

func TestCheckRule_1(t *testing.T) {
	tests := []struct {
		name    string
		rule    *Tree
		wantErr bool
	}{
		{
			name: "Test case 1: Rule with no args",
			rule: &Tree{
				Matcher: "matcher1",
				Value:   []string{},
			},
			wantErr: true,
		},
		{
			name: "Test case 2: Rule with empty args",
			rule: &Tree{
				Matcher: "matcher2",
				Value:   []string{""},
			},
			wantErr: true,
		},
		{
			name: "Test case 3: Rule with valid args",
			rule: &Tree{
				Matcher: "matcher3",
				Value:   []string{"arg1", "arg2"},
			},
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

			err := CheckRule(tt.rule)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
