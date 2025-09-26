package legacy

import (
	"fmt"
	"testing"

	"gopkg.in/ini.v1"
)

func TestCopySection_1(t *testing.T) {
	type args struct {
		source *ini.Section
		target *ini.Section
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

			copySection(tt.args.source, tt.args.target)
		})
	}
}
