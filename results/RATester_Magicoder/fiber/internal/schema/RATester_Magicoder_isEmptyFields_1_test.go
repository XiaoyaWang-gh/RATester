package schema

import (
	"fmt"
	"testing"
)

func TestisEmptyFields_1(t *testing.T) {
	type args struct {
		fields []fieldWithPrefix
		src    map[string][]string
	}
	tests := []struct {
		name string
		args args
		want bool
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

			if got := isEmptyFields(tt.args.fields, tt.args.src); got != tt.want {
				t.Errorf("isEmptyFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
