package hugolib

import (
	"fmt"
	"testing"
)

func TesthandleDuplicateResourcePath_1(t *testing.T) {
	type args struct {
		s        string
		updated  contentNodeI
		existing contentNodeI
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

			m := &pageMap{}
			m.handleDuplicateResourcePath(tt.args.s, tt.args.updated, tt.args.existing)
		})
	}
}
