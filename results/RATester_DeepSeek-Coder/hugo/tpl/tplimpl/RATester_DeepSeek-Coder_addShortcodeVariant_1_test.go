package tplimpl

import (
	"fmt"
	"testing"
)

func TestAddShortcodeVariant_1(t *testing.T) {
	type args struct {
		ts *templateState
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Errorf("templateHandler.addShortcodeVariant()")
		})
	}
}
