package context

import (
	"fmt"
	"testing"
)

func TestURI_1(t *testing.T) {
	type fields struct {
		Context *Context
	}
	tests := []struct {
		name   string
		fields fields
		want   string
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

			input := &BeegoInput{
				Context: tt.fields.Context,
			}
			if got := input.URI(); got != tt.want {
				t.Errorf("BeegoInput.URI() = %v, want %v", got, tt.want)
			}
		})
	}
}
