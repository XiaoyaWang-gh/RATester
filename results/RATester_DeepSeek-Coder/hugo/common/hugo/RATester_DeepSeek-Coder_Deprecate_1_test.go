package hugo

import (
	"fmt"
	"testing"
)

func TestDeprecate_1(t *testing.T) {
	type args struct {
		item        string
		alternative string
		version     string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestDeprecate",
			args: args{
				item:        "item",
				alternative: "alternative",
				version:     "version",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			Deprecate(tt.args.item, tt.args.alternative, tt.args.version)
		})
	}
}
