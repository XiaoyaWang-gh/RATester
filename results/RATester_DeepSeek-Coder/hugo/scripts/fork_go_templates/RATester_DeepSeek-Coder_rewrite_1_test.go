package main

import (
	"fmt"
	"testing"
)

func TestRewrite_1(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		rule     string
		wantErr  bool
	}{
		{
			name:     "success",
			filename: "test.go",
			rule:     "a:b",
			wantErr:  false,
		},
		{
			name:     "failure",
			filename: "test.go",
			rule:     "a:b",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); r != nil {
					if !tc.wantErr {
						t.Errorf("unexpected panic: %v", r)
					}
				}
			}()

			rewrite(tc.filename, tc.rule)
		})
	}
}
