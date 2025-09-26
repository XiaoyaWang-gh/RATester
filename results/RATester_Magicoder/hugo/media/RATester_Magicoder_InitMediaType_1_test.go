package media

import (
	"fmt"
	"testing"
)

func TestInitMediaType_1(t *testing.T) {
	type args struct {
		m *Type
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

			InitMediaType(tt.args.m)
		})
	}
}
