package commands

import (
	"fmt"
	"testing"
)

func TestInitMemProfile_1(t *testing.T) {
	type fields struct {
		r *rootCommand
	}
	tests := []struct {
		name   string
		fields fields
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

			c := &hugoBuilder{
				r: tt.fields.r,
			}
			c.initMemProfile()
		})
	}
}
