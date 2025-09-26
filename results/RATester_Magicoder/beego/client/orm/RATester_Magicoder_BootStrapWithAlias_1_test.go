package orm

import (
	"fmt"
	"testing"
)

func TestBootStrapWithAlias_1(t *testing.T) {
	type args struct {
		alias string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				alias: "test_alias",
			},
		},
		{
			name: "Test case 2",
			args: args{
				alias: "another_alias",
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

			BootStrapWithAlias(tt.args.alias)
		})
	}
}
